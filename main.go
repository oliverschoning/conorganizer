package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Regncon/conorganizer/models"
	"github.com/Regncon/conorganizer/routes"
	"github.com/Regncon/conorganizer/service"
	"github.com/go-chi/chi/v5"
	datastar "github.com/starfederation/datastar/sdk/go"
)

func main() {
	log.Println("Starting Regncon 2025")
	db, err := service.InitDB("events.db")
	if err != nil {
		log.Fatalf("Could not initialize DB: %v", err)
	}
	defer db.Close()

	router := chi.NewRouter()
	router.Use(loggingMiddleware)
	router.Use(recoveryMiddleware)

	router.Get("/", routes.RootRoute(db))
	router.Get("/event", routes.EventRoute())

	router.Route("/edit", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			var signals models.EditEvent
			err := datastar.ReadSignals(r, &signals)

			if err != nil {
				http.Error(w, fmt.Sprintf("Error reading signals: %v", err), http.StatusBadRequest)
				return
			}

			fmt.Printf("%+v signals\n", signals)
			sse := datastar.NewSSE(w, r)
			if err := sse.MergeSignals([]byte(fmt.Sprintf("%v", signals))); err != nil {
				http.Error(w, fmt.Sprintf("Error reading signals: %v", err), http.StatusBadRequest)
			}

		})

		r.Get("/{id}", routes.EditEventRoute(db))
	})

	fileServer := http.FileServer(http.Dir("./static"))
	router.Mount("/static", http.StripPrefix("/static/", fileServer))

	log.Printf("Listening on :3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Printf("error listening: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

// Recovery middleware to recover from panics
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
