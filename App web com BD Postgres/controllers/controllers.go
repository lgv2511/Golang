package controllers

import (
	"html/template"
	"fmt"
	"net/http"
	"log"
	"strconv"
	"alura3/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Obtendo lista de produtos...")
	todosOsProdutos := models.BuscaTodosProdutos()
	fmt.Println("Produtos recebidos do DB :", todosOsProdutos)
	fmt.Println("Executando template...")
	temp.ExecuteTemplate(w, "index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome  := r.FormValue("nome")
		descr := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quant := r.FormValue("quantidade")

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao obter preço do produto: ",err)
		}
		quantInt,err := strconv.Atoi(quant)
		if err != nil {
			log.Println("Erro ao obter quantidade do produto: ",err)
		}
		
		models.CriarProduto(nome, descr, quantInt, precoFloat)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProd := r.URL.Query().Get("id")
	produto := models.EditaProduto(idProd)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descr := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quant := r.FormValue("quantidade")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na convesão do preço para float64:", err)
		}

		quantInt, err := strconv.Atoi(quant)
		if err != nil {
			log.Println("Erro na convesão da quantidade para int:", err)
		}

		models.AtualizaProduto(idInt, nome, descr, precoFloat, quantInt)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
