package site

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/Jeffail/gabs/v2"
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	datastar "github.com/starfederation/datastar/sdk/go"
)

//	func test(w http.ResponseWriter, r *http.Request) {
//		sse := datastar.NewSSE(w, r)
//		sse.MergeFragments(
//			`<div id="question">What do you put in a toaster?</div>`,
//		)
//		sse.MergeSignals([]byte(`{response: '', answer: 'bread'}`))
//	}
func SetupExamplesTemplCounter(examplesRouter chi.Router, sessionStore sessions.Store) error {
	var globalCounter atomic.Uint32
	const (
		sessionKey = "templ_counter"
		countKey   = "count"
	)

	userVal := func(r *http.Request) (uint32, *sessions.Session, error) {
		sess, err := sessionStore.Get(r, sessionKey)
		if err != nil {
			return 0, nil, err
		}

		val, ok := sess.Values[countKey].(uint32)
		if !ok {
			val = 0
		}
		return val, sess, nil
	}

	examplesRouter.Get("/templ_counter/data", func(w http.ResponseWriter, r *http.Request) {
		userVal, _, err := userVal(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		store := TemplCounterStore{
			Global: globalCounter.Load(),
			User:   userVal,
		}

		sse := datastar.NewSSE(w, r)
		sse.MergeFragmentTempl(templCounterExampleInitialContents(store))
	})

	updateGlobal := func(store *gabs.Container) {
		store.Set(globalCounter.Add(1), "global")
	}

	examplesRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(templCounterExampleInitialContents2(TemplCounterStore{
			Global: 0,
			User:   0,
		})).Component.Render(r.Context(), w)

		examplesRouter.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			sse := datastar.NewSSE(w, r)
			sse.MergeFragments(
				`<div id="question">What do you put in a toaster?</div>`,
			)
			sse.MergeSignals([]byte(`{response: '', answer: 'bread'}`))
		})
	})

	examplesRouter.Route("/templ_counter/increment", func(incrementRouter chi.Router) {
		incrementRouter.Post("/global", func(w http.ResponseWriter, r *http.Request) {
			update := gabs.New()
			updateGlobal(update)

			fmt.Println(update.String())
			var signals TemplCounterStore
			err := datastar.ReadSignals(r, &signals)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error reading signals: %v", err), http.StatusBadRequest)
				return
			}
			fmt.Printf("%+v signals\n", signals)
			sse := datastar.NewSSE(w, r)
			sse.MergeSignals(update.Bytes())
		})

		incrementRouter.Post("/user", func(w http.ResponseWriter, r *http.Request) {
			val, sess, err := userVal(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			val++
			sess.Values[countKey] = val
			if err := sess.Save(r, w); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			update := gabs.New()
			updateGlobal(update)
			update.Set(val, "user")

			sse := datastar.NewSSE(w, r)
			sse.MergeSignals(update.Bytes())
		})
	})

	return nil
}
