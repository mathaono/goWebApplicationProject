package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Index", nil)
}
