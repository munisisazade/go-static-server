package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	//http.HandleFunc("/", handler) // each request calls handler
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveTemplate)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	lp := filepath.Join("index.html")
	tmpl, _ := template.ParseFiles(lp)
	tmpl.ExecuteTemplate(w, "index.html", "Hello")
}