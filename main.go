package main

import (
	"net/http"

	"github.com/goWebApplicationProject/routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":9000", nil)
}
