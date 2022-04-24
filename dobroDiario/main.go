//dobroDiario recebe um valor inicial e dobra-o a cada período por um total de 30 dias (ciclos).
package main

import "fmt"

func main() {
	var valor float64

	fmt.Println("Este programa recebe um valor inicial e o dobra a cada dia, ")
	fmt.Println("por um ciclo de 30 dias.")
	fmt.Print("Informe valor inicial: ")
	fmt.Scanf("%f", &valor)

	for i := 1; i <= 30; i++ {
		fmt.Printf("Dia %d - valor : %.2f\n", i, valor)
		valor = valor * 2
	}
}
