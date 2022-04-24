//Implementa trabalho proposto na formação básica em Golang da Alura.
//Faz o monitoramento de status on-line de alguns sites e gera log dos mesmos.

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const ciclos = 3
const tempo = 5

//menuOpções obtem resposta ao menu de opções do sistema.
func menuOpcoes() int {
	var retorno int

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Log")
	fmt.Println("3 - Encerrar Programa")
	fmt.Print("Escolha : ")

	fmt.Scanf("%d", &retorno)
	return retorno
}

//iniciaMonitoramento percorre os ciclos de monitoramento dos sites.
func iniciaMonitoramento() {
	fmt.Println("Iniciando monitoramento de sites... ")
	sites := obtemSites()

	for i := 1; i <= ciclos; i++ {
		fmt.Println("Ciclo de testes ", i, " de ", ciclos)
		for _, v := range sites {
			testaSite(v)
		}
		if i < ciclos {
			fmt.Println("Aguardando ", tempo, " segundos para próximo ciclo de testes")
			time.Sleep(tempo * time.Second)
		}
		fmt.Println("")
	}
}

//obtemSites recupera do arquivo 'sites.txt' a relação de sites a serem monitorados.
func obtemSites() []string {
	var slc []string

	arq, _ := os.Open("sites.txt")
	defer arq.Close()
	leitor := bufio.NewReader(arq)

	for {
		linhaarq, err := leitor.ReadString('\n')
		linhaarq = strings.TrimSpace(linhaarq)
		if err == io.EOF {
			break
		} else {
			slc = append(slc, linhaarq)
		}
	}
	return slc
}

//testaSite envia requisição http para site e registra resposta obtida no arquivo de log.
func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Falha na requisição de status do site ", site)
	} else {
		if resp.StatusCode == 200 {
			fmt.Println("Site ", site, " está on-line no momento")
			geraLog(site, true)
		} else {
			fmt.Println("Site ", site, " indisponível no momento - erro:", resp.StatusCode)
			geraLog(site, false)
		}
	}
}

//geraLog formata o registro de log e executa procedimentos de i/o com arquivo de log
func geraLog(site string, status bool) {
	arq, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	defer arq.Close()

	if err != nil {
		fmt.Println("Erro ao abrir arquivo de log - o mesmo não será registrado nesse ciclo de testes")
	} else {
		arq.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " +
			strconv.FormatBool(status) + "\n")
	}
}

//exibeLog recupera arquivo de log e exibe em tela.
func exibeLog() {
	arq, err := ioutil.ReadFile("log.txt")
	if err == nil {
		fmt.Println("Exibindo log de execução...")
		fmt.Println(string(arq))
	} else {
		fmt.Println("Erro ao tentar acesso ao log : ", err)
	}
}

func main() {

	for {
		comando := menuOpcoes()
		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			exibeLog()
		case 3:
			fmt.Println("Encerrando...")
			os.Exit(0)
		}
	}
}
