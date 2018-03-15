package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

// Content smthng
type Content struct {
	Txt string
}

// Wrapper smthng
type Wrapper struct {
	Title   string
	Content []Content
}

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

var templates = template.Must(template.ParseFiles(
	basepath+"/../template/index.html",
	basepath+"/../template/content.html",
	basepath+"/../template/login.html"))

func renderIndex(w http.ResponseWriter, r *http.Request) {
	handleError(w, templates.ExecuteTemplate(w, "index", Wrapper{"RRR", []Content{Content{Txt: "aaaa"}, Content{Txt: "bbb"}}}))
}

func renderLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("AAA" + r.FormValue("password"))
	log.Println("BBB" + r.FormValue("user_name"))
	handleError(w, templates.ExecuteTemplate(w, "login", nil))
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(basepath+"/../template/css"))))
	http.HandleFunc("/index", renderIndex)
	http.HandleFunc("/login", renderLogin)
	log.Println("sadfasdfa " + b)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
