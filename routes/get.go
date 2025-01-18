package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Regncon/conorganizer/components"
	"github.com/Regncon/conorganizer/models"
	"github.com/Regncon/conorganizer/pages/root"
	"github.com/a-h/templ"
	datastar "github.com/starfederation/datastar/sdk/go"
)

func RootRoute(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(root.Page(db)).Component.Render(r.Context(), w)
	}

}

func EventRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var signals models.EditEvent
		err := datastar.ReadSignals(r, &signals)
		if err != nil {
			// http.Error(w, fmt.Sprintf("Error reading signals: %v", err), http.StatusBadRequest)
			fmt.Printf("Accessed trough browser not signal %v", err)

		}
		event := models.Event{
			ID:          signals.ID,
			Name:        signals.Title,
			Description: signals.ShortDescription,
		}
		fmt.Printf("%+v signals\n", signals)
		sse := datastar.NewSSE(w, r)
		if err := sse.MergeFragmentTempl(root.EventCard(event, signals.System, signals.GameMaster)); err != nil {
			http.Error(w, fmt.Sprintf("Error reading signals: %v", err), http.StatusBadRequest)
		}

	}
}

func EditEventRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var signals models.EditEvent
		err := datastar.ReadSignals(r, &signals)
		if err != nil {
			// http.Error(w, fmt.Sprintf("Error reading signals: %v", err), http.StatusBadRequest)
			fmt.Printf("Accessed trough browser not signal %v", err)

		}
		fmt.Printf("%+v signals\n", signals)
		sse := datastar.NewSSE(w, r)
		if err := sse.MergeFragmentTempl(components.EditEvent(signals)); err != nil {
			http.Error(w, fmt.Sprintf("Error reading signals: %v", err), http.StatusBadRequest)
		}

	}
}
