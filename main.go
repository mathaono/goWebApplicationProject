package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
	Qtd       int
}

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}

func ConectDataBase() *sql.DB {
	connect := "user=postgres dbname=alura_loja password=adv17667 host=localhost sslmode=disable"
	db, error := sql.Open("postgres", connect)
	if error != nil {
		panic(error.Error())
	}
	return db
}

func Index(w http.ResponseWriter, r *http.Request) {
	db := ConectDataBase()

	selectProdutos, error := db.Query("select * from produtos")
	if error != nil {
		panic(error.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, qtd int
		var nome, descricao string
		var preco float64

		error = selectProdutos.Scan(&id, &nome, &descricao, &preco, &qtd)
		if error != nil {
			panic(error.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Qtd = qtd

		produtos = append(produtos, p)
	}
	tmpl.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
