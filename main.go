package main

import (
	"fmt"
	"net/http"
	"html/template"
	
	"database/sql"
	_ "github.com/lib/pq"
	
	"encoding/json"
)

type Page struct {
	Name string
	DBStatus bool
}

type SearchResult struct {
	Name string
	Description string
	Email string
	ID int
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))

	db, _ := sql.Open("postgres", "host=192.168.0.7 port=5432 user=postgres password=admin123 dbname=postgres sslmode=disable")
	// db, _ := sql.Open("postgres", "postgres://postgres:admin123@localhost/")
	_ = "breakpoint"
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		p := Page{Name: "Firuze"}
		
		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}
		p.DBStatus = db.Ping() == nil
		
		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// db.Close()
	})
	
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request){
		results := []SearchResult{
			SearchResult{"Jojon", "Standard user", "jojon@gmail.com", 1},
			SearchResult{"Sruntul", "Premium user", "sruntul@gmail.com", 2},
			SearchResult{"Cilok", "Pertamax user", "cilok@gmail.com", 3},
		}
		
		encoder := json.NewEncoder(w)
		if err := encoder.Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	
  fmt.Println(http.ListenAndServe(":8080", nil))
}