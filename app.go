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

	//// Return a 404 if the template doesn't exist
	//info, err := os.Stat(fp)
	//if err != nil {
	//	if os.IsNotExist(err) {
	//		http.NotFound(w, r)
	//		log.Println("Not exists")
	//		return
	//	}
	//}
	//
	//// Return a 404 if the request is for a directory
	//if info.IsDir() {
	//	http.NotFound(w, r)
	//	log.Println("Not Found")
	//	return
	//}
	//
	//tmpl, err := template.ParseFiles(lp, fp)
	//if err != nil {
	//	// Log the detailed error
	//	log.Println(err.Error())
	//	// Return a generic "Internal Server Error" message
	//	http.Error(w, http.StatusText(500), 500)
	//	log.Println("Internal Server Error")
	//	return
	//}
	//
	//if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
	//	log.Println(err.Error())
	//	http.Error(w, http.StatusText(500), 500)
	//	log.Println("Error")
	//}
}