//Main identifica números palíndromos
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func leNumero() int64 {
	var num int64

	fmt.Print("Informe número a ser analisado ou '0' para encerrar: ")
	fmt.Scanf("%d", &num)
	fmt.Scanf("%d", &num)
	if num == 0 {
		os.Exit(0)
	}
	return num
}

//Transforma em string o int64, inverte e compara, retorna 0 se iguais.
func analisaPalindromo(num int64) int {
	var strInv string

	numStr := strconv.FormatInt(num, 10)
	for i := len(numStr) - 1; i >= 0; i-- {
		strInv += string(numStr[i])
	}
	fmt.Println("Original: ", numStr, " - Invertido: ", strInv)
	return strings.Compare(numStr, strInv)
}

func main() {
	fmt.Println("Um número é chamado palíndromo quando lido de trás pra frente é igual de frente para trás.")
	for {
		num := leNumero()
		if num > 0 && (num%10 != 0) {
			if analisaPalindromo(num) == 0 {
				fmt.Println("O número ", num, " é palíndromo")
			} else {
				fmt.Println("O número ", num, " NAO é palíndromo")
			}
		} else {
			fmt.Println("Números negativos ou que terminam em zero não são palíndromos")
		}
	}
}
