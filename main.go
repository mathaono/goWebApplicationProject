package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Nome      string
	Descricao string
	Preco     float64
	Qtd       int
}

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{"Camiseta", "Vermelha", 59.90, 10},
		{"Calça Jeans", "Preta", 119.90, 8},
		{"Tênis", "Azul", 199.90, 5},
		{"Blusa Moletom", "Branca", 189.90, 20},
	}

	tmpl.ExecuteTemplate(w, "Index", produtos)
}
