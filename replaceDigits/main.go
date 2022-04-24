//ReplaceDigits implementa o clássico algoritmo de substituir todos os dígitos por caracteres.
//Utiliza arquivo 'cadeias.txt' com as sequências de caracteres a serem convertidas.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const arqtxt = "cadeias.txt"

//ObtemCadeiasCaracteres retorna um slice com as cadeias encontradas no arquivo texto 'cadeias.txt'
func obtemCadeiasCaracteres() []string {
	var slc []string

	arq, _ := os.Open(arqtxt)
	defer arq.Close()
	leitor := bufio.NewReader(arq)

	for {
		linhaarq, err := leitor.ReadString('\n')
		linhaarq = strings.TrimSpace(linhaarq)
		if err == io.EOF {
			return slc
		} else {
			slc = append(slc, linhaarq)
		}
	}
}

//Converte percorre o parâmetro string convertendo dígitos por caracteres e retorna.
func Converte(inicial string) string {
	var saida string
	for i := 1; i < len(inicial); i = i + 2 {
		saida += string(inicial[i-1])
		pos, _ := strconv.Atoi(string(inicial[i]))
		saida += string(int(inicial[i-1]) + pos)
	}
	if len(inicial)%2 != 0 {
		saida += string(inicial[len(inicial)-1])
	}
	return saida
}

func main() {
	cadeias := obtemCadeiasCaracteres()

	for i, v := range cadeias {
		result := Converte(v)
		fmt.Println("Convertendo digitos:", cadeias[i], "\t>> Resultado:", result)
	}
}
