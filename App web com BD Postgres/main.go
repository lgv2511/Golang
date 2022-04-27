package main

import (
	"alura3/routes"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Chamando Carrega Rotas")
	routes.CarregaRotas()
	fmt.Println("Escaneando porta 8000...")
	http.ListenAndServe(":8000", nil)
}
