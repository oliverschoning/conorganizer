package routes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Regncon/conorganizer/models"
	"github.com/Regncon/conorganizer/pages/root"
	datastar "github.com/starfederation/datastar/sdk/go"
)

func UpdateEventRoute(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// defer rows.Close()

		// var events []models.Event
		// for rows.Next() {
		// 	var event models.Event
		// 	if err := rows.Scan(&event.ID, &event.Name, &event.Description); err != nil {
		// 		return nil, err
		// 	}
		// 	events = append(events, event)
		// }
		// return events, nil

		var signals models.Event
		err := datastar.ReadSignals(r, &signals)
		if err != nil {
			// http.Error(w, fmt.Sprintf("Error reading signals: %v", err), http.StatusBadRequest)
			fmt.Printf("Accessed trough browser not signal %v", err)

		}

		query := "UPDATE events SET title = ?, short_description = ?, game_master = ?, system = ? WHERE id = ?"
		res, err := db.Exec(query, signals.Title, signals.ShortDescription, signals.GameMaster, signals.System, signals.ID)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error updating event: %v", err), http.StatusBadRequest)
			return
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error fetching rows affected: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if rowsAffected == 0 {
			http.Error(w, "Event not found or no changes made", http.StatusNotFound)
			return
		}

		fmt.Printf("%+v signals\n", signals)
		sse := datastar.NewSSE(w, r)
		if err := sse.MergeFragmentTempl(root.EventCard(signals, int(signals.ID))); err != nil {
			http.Error(w, fmt.Sprintf("Error reading signals: %v", err), http.StatusBadRequest)
		}

	}
}
