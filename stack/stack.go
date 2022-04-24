package main

import "fmt"

type Item struct {
	Last  int
	Value string
}

var Stack []Item

func Insere(valor string) {
	var novo Item

	novo.Last = len(Stack) - 1
	novo.Value = valor
	Stack = append(Stack, novo)

	fmt.Println("Inserindo na posição", novo.Last+1, "item", valor)
}

func MostraPilha() {
	fmt.Print("Pilha: <início> ")

	for i := 0; i < len(Stack); i++ {
		fmt.Print(Stack[i].Value, " - ")
	}
	fmt.Println("<final>")
}

func Remove() {
	fmt.Println("Retirando último elemento: ", Stack[len(Stack)-1].Value)

	x := len(Stack) - 1
	Stack = append(Stack[:x])
	MostraPilha()
}
