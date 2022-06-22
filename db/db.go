package db

//CÓDIGO RESPONSÁVEL PELA CONEXÃO E CONFIGURAÇÃO DO BANCO DE DADOS
import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	//Armazenado na variável todas as configs do banco (usuário, senha, e endereçamento do banco)
	conexao := "user=postgres dbname=alura_loja password=adv17667 host=localhost sslmode=disable"

	//Abrindo conexão com o banco
	db, err := sql.Open("postgres", conexao)

	//Validação de erro
	if err != nil {
		panic(err.Error())
	}

	//Retornando conexão com o banco
	//Ao invés de executar esse bloco de código em todo lugar, basta importar o pacote db e chamar a função ConectaComBancoDeDados
	return db
}
