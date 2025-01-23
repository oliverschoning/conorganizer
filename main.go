package main

import (
	"log"
	"net/http"
	"html/template"

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
		err = tmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// start server, log errors
	log.Fatal(http.ListenAndServe(":8080", nil))
}
