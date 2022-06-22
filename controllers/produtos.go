package controllers

//CÓDIO RESPONSÁVEL PELO CONTROLE DE EXECUÇÃO DA APLICAÇÃO
import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/goWebApplicationProject/models"
)

//Armazenando em uma variável uma função que executa as páginas no diretório /templates; No caso, todos os arquivos .html são chamados
var tmpl = template.Must(template.ParseGlob("templates/*.html"))

//Funcção p/ chamar a página inicial da aplicação, no caso, a página index.html que é chamada na url "/" configurado em routes.go
func Index(w http.ResponseWriter, r *http.Request) {

	//Chamando a função do diretório models p/ pesquisa de todos os produtos
	todosOsProdutos := models.BuscaTodosOsProdutos()
	tmpl.ExecuteTemplate(w, "Index", todosOsProdutos)
}

//Função apenas p/ executar um novo template, no caso, o arquivo new.html que é um template de formulário, que será usado p/ add novos produtos
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

//Função p/ add novo produto
func Insert(w http.ResponseWriter, r *http.Request) {

	//Verificando se a requisição é um POST
	if r.Method == "POST" {

		//Se atender a condição, armazena em variáveis todos os campos necessários p/ add novo produto ao banco
		//Os nomes dos campos são adicionados no próprio html com o parâmetro "id=idDoCampo"
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		qtd := r.FormValue("qtd")

		//Conversão p/ float64
		precoConvertido, error := strconv.ParseFloat(preco, 64)
		if error != nil {
			fmt.Println("Erro na conversão do preco: ", error)
		}

		//Conversão p/ inteiro
		qtdConvertida, error := strconv.Atoi(qtd)
		if error != nil {
			fmt.Println("Erro na conversão da quantidade: ", error)
		}

		//Chamando a função de criação de novo produto do diretório models e passando os respectivos campos como parâmetro
		models.CriaNovoProduto(nome, descricao, precoConvertido, qtdConvertida)
	}
	//Redirecionando p/ a tela principal após add produto novo
	http.Redirect(w, r, "/", 301)
}

//Função p/ deletar produto do banco
func Delete(w http.ResponseWriter, r *http.Request) {

	//Salvando o id do produto em uma variável; Esse Id é o primary key (auto incrementado) do banco de dados
	idProduto := r.URL.Query().Get("id")

	//Chamando a função de exclusão do produto no banco e passando como parâmetro, a variável com o Id armazenado
	models.DeletaProduto(idProduto)

	//Redirecionando p/ a tela principal após excluir o produto
	http.Redirect(w, r, "/", 301)
}
