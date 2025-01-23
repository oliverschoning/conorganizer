package main

import (
	"log"
	"net/http"
	"html/template"
	"encoding/json"

// todo: mux for dynamic routing, it has been installed,
	//"github.com/gorilla/mux"
)

func main() {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		log.Fatal(err)
	}
	// simple handler to demonstrate
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		data := struct {
			Title string
		}{
			Title: "Regncon: program",
		}
		err = tmpl.ExecuteTemplate(w, "layout.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// start server, log errors
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// sample data
	data := map[string]string{"message":"Hello from HTMX"}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().set("Content-Type", "application/json")

	
}