package index

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/delaneyj/toolbelt"
	"github.com/delaneyj/toolbelt/embeddednats"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/nats-io/nats.go/jetstream"
	datastar "github.com/starfederation/datastar/sdk/go"
)

func SetupIndexRoute(router chi.Router, store sessions.Store, ns *embeddednats.Server, db *sql.DB, logger *slog.Logger) error {
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

	saveMVC := func(ctx context.Context, sessionID string, mvc *EventState) error {
		b, err := json.Marshal(mvc)
		if err != nil {
			return fmt.Errorf("failed to marshal mvc: %w", err)
		}
		if _, err := kv.Put(ctx, sessionID, b); err != nil {
			return fmt.Errorf("failed to put key value: %w", err)
		}
		return nil
	}

	resetMVC := func(mvc *EventState) {
		events, err := getEvents(db, logger)
		if err != nil {
			logger.Error("Could not get events", "err", err)
		}

		mvc.Events = events

		mvc.Idx = -1
	}

	mvcSession := func(w http.ResponseWriter, r *http.Request) (string, *EventState, error) {
		ctx := r.Context()
		sessionID, err := upsertSessionID(store, r, w)
		if err != nil {
			return "", nil, fmt.Errorf("failed to get session id: %w", err)
		}

		mvc := &EventState{}
		if entry, err := kv.Get(ctx, sessionID); err != nil {
			if err != jetstream.ErrKeyNotFound {
				return "", nil, fmt.Errorf("failed to get key value: %w", err)
			}
			resetMVC(mvc)

			if err := saveMVC(ctx, sessionID, mvc); err != nil {
				return "", nil, fmt.Errorf("failed to save mvc: %w", err)
			}
		} else {
			if err := json.Unmarshal(entry.Value(), mvc); err != nil {
				return "", nil, fmt.Errorf("failed to unmarshal mvc: %w", err)
			}
		}
		return sessionID, mvc, nil
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Index("Regncon program 2025", "/events/stream").Render(r.Context(), w)
	})

	router.Route("/events/stream", func(eventsRouter chi.Router) {
		eventsRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
			sessionID, mvc, err := mvcSession(w, r)
			if err != nil {
				logger.Error("failed getting mvc session", "err", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			logger.Info("Session ID and MVC retrieved successfully", "sessionID", sessionID)

			sse := datastar.NewSSE(w, r)

			// Watch for updates
			ctx := r.Context()
			watcher, err := kv.Watch(ctx, sessionID)
			if err != nil {
				logger.Error("Error creating watcher", "err", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer watcher.Stop()
			logger.Info("Watcher created successfully")

			for {
				select {
				case <-ctx.Done():
					logger.Info("Context done")
					return
				case entry := <-watcher.Updates():
					if entry == nil {
						continue
					}
					if err := json.Unmarshal(entry.Value(), mvc); err != nil {
						logger.Error("Error unmarshalling entry", "err", err)
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					c := EventsView(mvc, db, logger)
					if err := sse.MergeFragmentTempl(c); err != nil {
						logger.Error("Error merging fragment template", "err", err)
						sse.ConsoleError(err)
						return
					}
				}
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
