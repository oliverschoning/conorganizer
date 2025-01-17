package main

import (
	"log"
	"net/http"
	"time"

	site "github.com/Regncon/conorganizer/test"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	// "github.com/Regncon/conorganizer/pages/event/add"
	// "github.com/Regncon/conorganizer/pages/event/edit"
	// "github.com/Regncon/conorganizer/pages/root"
	// "github.com/Regncon/conorganizer/service"
	// "github.com/a-h/templ"
)

func main() {
	// log.Println("Starting Regncon 2025")
	// db, err := service.InitDB("events.db")
	// if err != nil {
	// 	log.Fatalf("Could not initialize DB: %v", err)
	// }
	// defer db.Close()

	// http.Handle("/", templ.Handler(root.Page(db)))
	// http.Handle("/event/add/", templ.Handler(add.Page()))
	// http.HandleFunc("/event/", func(w http.ResponseWriter, r *http.Request) {
	// 	// Extract the event ID from the URL
	// 	path := strings.TrimPrefix(r.URL.Path, "/event/")
	// 	if path == "" || strings.Contains(path, "/") {
	// 		http.NotFound(w, r)
	// 		return
	// 	}

	// 	log.Printf("Event handler for ID: %s", path)
	// 	// Call the event page handler with the extracted ID
	// 	templ.Handler(edit.Page(path, db)).Component.Render(r.Context(), w)
	// })

	// http.HandleFunc("/event/edit/save/", edit.Save(db))

	// http.HandleFunc("/event/add/new/", func(w http.ResponseWriter, r *http.Request) {
	// 	templ.Handler(add.EventNew(w, r, db)).Component.Render(r.Context(), w)
	// })

	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle POST and GET requests.

	sessionStore := sessions.NewCookieStore([]byte("your-secret-key"))

	r := chi.NewRouter()
	r.Use(loggingMiddleware)
	r.Use(recoveryMiddleware)

	err := site.SetupExamplesTemplCounter(r, sessionStore)
	if err != nil {
		log.Fatalf("Failed to set up templ counter: %v", err)
	}

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	log.Println("Listening on :3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
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
