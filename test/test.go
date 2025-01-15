package site

import (
	"net/http"
	"sync/atomic"

	"github.com/Jeffail/gabs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	datastar "github.com/starfederation/datastar/sdk/go"
)

type TemplCounterStore struct {
	Global uint32 `json:"global"`
	User   uint32 `json:"user"`
}

templ templCounterExampleButtons() {
	<div>
		<button
			data-on-click="$$post('/examples/templ_counter/increment/global')"
		>
			Increment Global
		</button>
		<button
			data-on-click="$$post('/examples/templ_counter/increment/user')"
		>
			Increment User
		</button>
	</div>
}

templ templCounterExampleCounts() {
	<div>
		<div>
			<div>Global</div>
			<div data-text="$global"></div>
		</div>
		<div>
			<div>User</div>
			<div data-text="$user"></div>
		</div>
	</div>
}

templ templCounterExampleInitialContents(store TemplCounterStore) {
	<div
		id="container"
		data-store={ templ.JSONString(store) }
	>
		@templCounterExampleButtons()
		@templCounterExampleCounts()
	</div>
}

func setupExamplesTemplCounter(examplesRouter chi.Router, sessionStore sessions.Store) error {
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
		datastar.RenderFragmentTempl(sse, templCounterExampleInitialContents(store))
	})

	updateGlobal := func(store *gabs.Container) {
		store.Set(globalCounter.Add(1), "global")
	}

	examplesRouter.Route("/templ_counter/increment", func(incrementRouter chi.Router) {
		incrementRouter.Post("/global", func(w http.ResponseWriter, r *http.Request) {
			update := gabs.New()
			updateGlobal(update)

			sse := datastar.NewSSE(w, r)
			datastar.PatchStore(sse, update)
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
			datastar.PatchStore(sse, update)
		})
	})

	return nil
}
