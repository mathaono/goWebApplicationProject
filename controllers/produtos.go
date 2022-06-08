package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/goWebApplicationProject/models"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaTodosOsProdutos()
	tmpl.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		qtd := r.FormValue("qtd")

		precoConvertido, error := strconv.ParseFloat(preco, 64)
		if error != nil {
			fmt.Println("Erro na conversão do preco: ", error)
		}

		qtdConvertida, error := strconv.Atoi(qtd)
		if error != nil {
			fmt.Println("Erro na conversão da quantidade: ", error)
		}

		models.CriaNovoProduto(nome, descricao, precoConvertido, qtdConvertida)
	}
	http.Redirect(w, r, "/", 301)
}
