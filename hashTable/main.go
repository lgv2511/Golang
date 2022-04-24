// Main demonstra uma lista organizada via hash code gerado com base no nome
package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Sexo      string
}

func main() {
	pessoas := []Pessoa{
		{"Ricardo", "Silva", 58, "M"},
		{"Sonia", "Santos", 60, "F"},
		{"Maria", "Almeida", 55, "F"},
		{"Nelson", "Pires", 70, "M"},
	}

	table := Hashtable{}
	keys := make([]int, len(pessoas))

	for i, pessoa := range pessoas {
		keys[i] = table.Put(pessoa)
		fmt.Println(" >> Inserindo", pessoa.Nome, pessoa.Sobrenome, "com a hash", keys[i])
	}

	fmt.Println("--------------\nLista atual:")
	for _, key := range keys {
		ps := table.Get(key)
		for _, p := range ps {
			fmt.Println(p.Nome, p.Sobrenome)
		}
	}

	table.Remove("Sonia")

	proc := table.Busca("Nelson")
	fmt.Println("Busca realizada com sucesso:", proc)
}
