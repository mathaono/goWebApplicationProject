package models

import "github.com/goWebApplicationProject/db"

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
	Qtd       int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, qtd int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &qtd)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Qtd = qtd

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, qtd int) {
	db := db.ConectaComBancoDeDados()

	insertNovoProduto, error := db.Prepare("insert into produtos (nome, descricao, preco, qtd) values ($1, $2, $3, $4)")
	if error != nil {
		panic(error.Error())
	}

	insertNovoProduto.Exec(nome, descricao, preco, qtd)
	defer db.Close()
}
