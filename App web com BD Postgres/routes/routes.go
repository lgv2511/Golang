package routes

import (
	"alura3/controllers"
	"net/http"
	"fmt"

	_ "github.com/lib/pq"
)

func CarregaRotas() {
	fmt.Println("Inicializando Rotas...")
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
