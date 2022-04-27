package models

import (
	"alura3/db"
	"fmt"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quant           int
}

func BuscaTodosProdutos() []Produto {
	fmt.Println("Abrindo banco de dados...")
	db := db.ConectaBancoDados()
	defer db.Close()

	selectProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Obtendo produtos no BD...")
	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quant int
		var nome, descricao string
		var preco float64
		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quant)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quant = quant

		produtos = append(produtos, p)
	}
	return produtos
}

func CriarProduto(nome, descr string, quant int, preco float64) {
	db := db.ConectaBancoDados()

	insereBD, err := db.Prepare("insert into produtos(nome,descricao,preco,quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insereBD.Exec(nome, descr, preco, quant)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaBancoDados()

	deletarProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarProduto.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConectaBancoDados()

	produtoDB, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDB.Next() {
		var id, quant int
		var nome, descr string
		var preco float64

		err = produtoDB.Scan(&id, &nome, &descr, &preco, &quant)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descr
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quant = quant
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, descr string, preco float64, quant int) {
	db := db.ConectaBancoDados()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descr, preco, quant, id)
	defer db.Close()
}
