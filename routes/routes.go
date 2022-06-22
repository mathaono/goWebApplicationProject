package routes

//CÓDIGO RESPONSÁVEL PELO DIRECIONAMENTO DE ROTAS DA NOSSA APLICAÇÃO
import (
	"net/http"

	"github.com/goWebApplicationProject/controllers"
)

//Carregamento de rotas e execução das funções nessas rotas
func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
}
