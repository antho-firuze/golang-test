package main

import (
	"fmt"
	"net/http"
	"html/template"
	
	"database/sql"
	_ "github.com/lib/pq"
)

type Page struct {
	Name string
	DBStatus bool
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
		db.Close()
	})
	
  fmt.Println(http.ListenAndServe(":8080", nil))
}