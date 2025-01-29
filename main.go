package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	// todo: mux for dynamic routing, it has been installed,
	// "github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatal("failed to open db", err)
	}
	if db == nil {
		log.Fatal("db is nil, something is clearly wrong")
	}
	defer db.Close()

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		log.Fatal(err)
	}
	// simple handler to demonstrate
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Title string
		}{
			Title: "Regncon: program",
		}
		w.Header().Set("Content-Type", "text/html")
		err = tmpl.ExecuteTemplate(w, "layout.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/data", dataHandler)

	// start server, log errors
	log.Println("Server starting on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func dataHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("SELECT * FROM rooms")
	if err != nil {
		log.Fatal("Failed to execute query", err)
	}

	var data string
	rows.Next()
	err = rows.Scan(&data)
	jsonData, err := json.Marshal(data)
	log.Println(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
