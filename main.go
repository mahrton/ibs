package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var templates = template.Must(template.ParseFiles("template/login.html"))

func createLoginRender(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("AAA" + r.FormValue("password"))
		log.Println("BBB" + r.FormValue("user_name"))
		handleError(w, templates.ExecuteTemplate(w, "login", nil))
	}
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	db, err := sql.Open("mysql", "ibs:admin@tcp(127.0.0.1:3306)/ibs_login")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("template/css"))))
	http.HandleFunc("/login", createLoginRender(db))
	log.Fatal(http.ListenAndServe(":3001", nil))
}
