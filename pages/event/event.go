package event

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/Regncon/conorganizer/models"
	"github.com/Regncon/conorganizer/pages/index"
	"github.com/delaneyj/toolbelt"
	"github.com/delaneyj/toolbelt/embeddednats"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/nats-io/nats.go/jetstream"
	datastar "github.com/starfederation/datastar/sdk/go"
)

func SetupEventRoute(router chi.Router, store sessions.Store, ns *embeddednats.Server, db *sql.DB, logger *slog.Logger) error {
	nc, err := ns.Client()
	if err != nil {
		return fmt.Errorf("error creating nats client: %w", err)
	}
	logger.Info("NATS client created successfully")
	js, err := jetstream.New(nc)
	if err != nil {
		return fmt.Errorf("error creating jetstream client: %w", err)
	}
	logger.Info("JetStream client created successfully")

	kv, err := js.CreateOrUpdateKeyValue(context.Background(), jetstream.KeyValueConfig{
		Bucket:      "regncon",
		Description: "regncon 2025",
		Compression: true,
		TTL:         time.Hour,
		MaxBytes:    16 * 1024 * 1024,
	})

	if err != nil {
		return fmt.Errorf("error creating key value: %w", err)
	}
	logger.Info("Key value store created successfully")

	saveState := func(ctx context.Context, sessionID string, state *EventState) error {
		b, err := json.Marshal(state)
		if err != nil {
			return fmt.Errorf("failed to marshal state: %w", err)
		}
		if _, err := kv.Put(ctx, sessionID, b); err != nil {
			return fmt.Errorf("failed to put key value: %w", err)
		}
		return nil
	}

	resetState := func(state *EventState, id int) {
		event, err := GetEvent(db, id, logger)
		if err != nil {
			logger.Error("Could not get events", "err", err)
		}

		state.Event = event
		state.Idx = id
	}

	stateSessionWithId := func(w http.ResponseWriter, r *http.Request, id int) (string, *EventState, error) {
		ctx := r.Context()
		sessionID, err := upsertSessionID(store, r, w)
		if err != nil {
			return "", nil, fmt.Errorf("failed to get session id: %w", err)
		}

		state := &EventState{}
		if entry, err := kv.Get(ctx, sessionID); err != nil {
			if err != jetstream.ErrKeyNotFound {
				return "", nil, fmt.Errorf("failed to get key value: %w", err)
			}
			resetState(state, id)

			if err := saveState(ctx, sessionID, state); err != nil {
				return "", nil, fmt.Errorf("failed to save state: %w", err)
			}
		} else {
			if err := json.Unmarshal(entry.Value(), state); err != nil {
				return "", nil, fmt.Errorf("failed to unmarshal state: %w", err)
			}
		}
		return sessionID, state, nil
	}

	router.Route("/event", func(eventRouter chi.Router) {
		// eventRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {

		// 	index.Index("Events", "/event").Render(r.Context(), w)
		// })
		eventRouter.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {

			idString := chi.URLParam(r, "id")
			id, err := strconv.Atoi(idString)
			if err != nil {
				logger.Error("Id not found", "err", err, "id", idString)
				http.Error(w, "ID must be numeric", http.StatusBadRequest)
				return
			}
			index.Index("Events", fmt.Sprintf("/event/%d", id)).Render(r.Context(), w)
			sessionID, state, err := stateSessionWithId(w, r, id)
			sse := datastar.NewSSE(w, r)
			if err != nil {
				sse.ConsoleError(err)
				return
			}

			state.Idx = id
			if err := saveState(r.Context(), sessionID, state); err != nil {
				sse.ConsoleError(err)
				return
			}

		})
		eventRouter.Put("/{id}/edit", func(w http.ResponseWriter, r *http.Request) {
			idString := chi.URLParam(r, "id")
			id, err := strconv.Atoi(idString)
			if err != nil {
				logger.Error("Id not found", "err", err, "id", idString)
				http.Error(w, "ID must be numeric", http.StatusBadRequest)
				return
			}

			sessionID, state, err := stateSessionWithId(w, r, id)
			sse := datastar.NewSSE(w, r)
			if err != nil {
				sse.ConsoleError(err)
				return
			}

			state.Idx = id
			state.IsEditing = true
			if err := saveState(r.Context(), sessionID, state); err != nil {
				sse.ConsoleError(err)
				return
			}

		})
		eventRouter.Put("/{id}/save", func(w http.ResponseWriter, r *http.Request) {
			idString := chi.URLParam(r, "id")
			id, err := strconv.Atoi(idString)
			if err != nil {
				logger.Error("Id not found", "err", err, "id", idString)
				http.Error(w, "ID must be numeric", http.StatusBadRequest)
				return
			}

			var updatedEvent models.Event
			updatedEvent.ID = int64(id)
			err = datastar.ReadSignals(r, &updatedEvent)
			if err != nil {
				logger.Error("Error reading signals", "err", err)
				fmt.Printf("Accessed trough browser not signal %v", err)
			}

			sessionID, state, err := stateSessionWithId(w, r, id)
			sse := datastar.NewSSE(w, r)
			if err != nil {
				sse.ConsoleError(err)
				return
			}

			err = updateDb(db, logger, updatedEvent, w, r)
			if err != nil {
				logger.Error("Error updating event", "err", err)
				sse.ConsoleError(err)
				return
			}
			state.Event = updatedEvent
			state.IsEditing = false

			if err := saveState(r.Context(), sessionID, state); err != nil {
				sse.ConsoleError(err)
				return
			}

		})
		eventRouter.Put("/cancel", func(w http.ResponseWriter, r *http.Request) {

			sessionID, state, err := stateSessionWithId(w, r, -1)
			sse := datastar.NewSSE(w, r)
			if err != nil {
				sse.ConsoleError(err)
				return
			}

			state.IsEditing = false
			if err := saveState(r.Context(), sessionID, state); err != nil {
				sse.ConsoleError(err)
				return
			}

		})
		eventRouter.Put("/back", func(w http.ResponseWriter, r *http.Request) {

			sessionID, state, err := stateSessionWithId(w, r, -1)
			sse := datastar.NewSSE(w, r)
			if err != nil {
				sse.ConsoleError(err)
				return
			}

			state.Idx = -1
			state.IsEditing = false
			if err := saveState(r.Context(), sessionID, state); err != nil {
				sse.ConsoleError(err)
				return
			}

		})
	})

	return nil
}

func upsertSessionID(store sessions.Store, r *http.Request, w http.ResponseWriter) (string, error) {

	sess, err := store.Get(r, "connections")
	if err != nil {
		return "", fmt.Errorf("failed to get session: %w", err)
	}
	id, ok := sess.Values["id"].(string)
	if !ok {
		id = toolbelt.NextEncodedID()
		sess.Values["id"] = id
		if err := sess.Save(r, w); err != nil {
			return "", fmt.Errorf("failed to save session: %w", err)
		}
	}
	return id, nil
}

func updateDb(db *sql.DB, logger *slog.Logger, updatedEvent models.Event, w http.ResponseWriter, r *http.Request) error {
	query := "UPDATE events SET title = ?, short_description = ?, game_master = ?, system = ? WHERE id = ?"
	res, err := db.Exec(query, updatedEvent.Title, updatedEvent.ShortDescription, updatedEvent.GameMaster, updatedEvent.System, updatedEvent.ID)
	if err != nil {
		logger.Error("Error updating event", "err", err)
		http.Error(w, fmt.Sprintf("Error updating event: %v", err), http.StatusBadRequest)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logger.Error("Error getting rows affected", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return err
	}

	if rowsAffected == 0 {
		logger.Error("Event not found or no changes made")
		http.Error(w, "Event not found or no changes made", http.StatusNotFound)
		return err
	}
	fmt.Printf("Event updated successfully: %v\n", updatedEvent)
	return nil
}
