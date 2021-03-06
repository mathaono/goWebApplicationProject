package models

//CÓDIGO RESPONSÁVEL PELAS AÇÕES NA NOSSA APLICAÇÃO EM RELAÇÃO AO BANCO DE DADOS
import "github.com/goWebApplicationProject/db"

//Struct responsável por capturar os dados dos produtos (novos e os já existentes)
type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
	Qtd       int
}

func BuscaTodosOsProdutos() []Produto {
	//Conexão com o banco
	db := db.ConectaComBancoDeDados()

	//Preparando a Query para pesquisa de produtos no banco de dados
	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	//Armazenando a struct em uma variável
	p := Produto{}

	//Armazenando e transformando a struct em uma lista (ou slice) de structs
	produtos := []Produto{}

	//Executando a query em um loop
	for selectDeTodosOsProdutos.Next() {
		var id, qtd int
		var nome, descricao string
		var preco float64

		//Escaneando todos os campos p/ o caso de erro
		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &qtd)
		if err != nil {
			panic(err.Error())
		}

		//Armazenando cada item do banco dentro da variável (p) que armazena a struct
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Qtd = qtd

		//Adicionando cada variável struct (p) dentro da lista de structs (produtos)
		produtos = append(produtos, p)
	}

	//Encerrando conexão e retornando a lista de structs
	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, qtd int) {
	//Abrindo conexão com o banco
	db := db.ConectaComBancoDeDados()

	//Preparando a query p/ adicionar novo item ao banco de dados
	insertNovoProduto, error := db.Prepare("insert into produtos (nome, descricao, preco, qtd) values ($1, $2, $3, $4)")
	if error != nil {
		panic(error.Error())
	}

	//Executando a query passando os parâmetros necessários p/ adição no banco
	insertNovoProduto.Exec(nome, descricao, preco, qtd)

	//Encerrando conexão com o banco
	defer db.Close()
}

func DeletaProduto(id string) {
	//Conexão com o banco
	db := db.ConectaComBancoDeDados()

	//Preparando a Query para deletar no banco de dados
	queryDeletaProduto, error := db.Prepare("delete from produtos where id=$1")
	if error != nil {
		panic(error.Error())
	}

	//Executando a query passando o parâmetro identificador do item a ser deletado
	queryDeletaProduto.Exec(id)

	//Encerrando conexão com banco
	defer db.Close()
}

//Função p/ editar um produto
func EditaProduto(id string) Produto {
	//Conexão com o banco
	db := db.ConectaComBancoDeDados()

	//Preparando a query
	produtoDoBanco, error := db.Query("select * from produtos where id=$1", id)
	if error != nil {
		panic(error.Error())
	}

	//Armazenando uma instância da struct de Produto em uma variável
	atualizaProduto := Produto{}

	//Usando um for junto à função Next, que percorre e trás os dados de uma linha na tabela do banco
	//Cada vez que o for é executado o Next() pega a linha seguinte
	for produtoDoBanco.Next() {
		var id, qtd int
		var nome, descricao string
		var preco float64

		//Usando scan p/ armazenar os respectivos campos do banco em suas respectivas variáveis declaradas acima
		error = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &qtd)
		if error != nil {
			panic(error.Error())
		}

		//Passando essas variáveis para a variável que armazena uma instância da struct Produto, armazenando cada variável em seu respectivo campo
		atualizaProduto.Id = id
		atualizaProduto.Nome = nome
		atualizaProduto.Descricao = descricao
		atualizaProduto.Qtd = qtd
		atualizaProduto.Preco = preco
	}
	defer db.Close()

	//Retornando a struct de produtos armazenada na variável
	return atualizaProduto
}

func AtualizaProduto(id int, nome, descricao string, preco float64, qtd int) {
	//Conexão com o banco
	db := db.ConectaComBancoDeDados()

	atualizandoProduto, error := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, qtd=$4 where id=$5")
	if error != nil {
		panic(error.Error())
	}

	atualizandoProduto.Exec(nome, descricao, preco, qtd, id)
	db.Close()
}
