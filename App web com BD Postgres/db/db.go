package db

import (
	"fmt"
	"database/sql"
	_"github.com/lib/pq"
)
	
func ConectaBancoDados() *sql.DB {
	fmt.Println("Estabelecendo conexão com servidor de banco de dados...")
	conexao := "user=postgres dbname=alura_loja password=lgvpost host=localhost sslmode=disable"
	bd, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error)
	}
	return bd
}
