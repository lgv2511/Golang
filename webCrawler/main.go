//WebCrawler implementa um rastreador básico de links dentro de uma página da internet.
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var (
	links   []string
	visited map[string]bool = map[string]bool{}
)

func menu() string {
	var ret string

	fmt.Print("Informe o site a ser analisado ou tecle enter para encerrar: ")
	fmt.Scanf("%s", &ret)
	return ret
}

func verificaLink(url string) {
	if ok := visited[url]; ok {
		return
	}

	visited[url] = true
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		msg := ">>>> O site " + url + " não está on-line no momento"
		geraLog(msg)
		fmt.Println(msg, " Encerrando execução...")
		os.Exit(0)
	} else {
		doc, err := html.Parse(resp.Body)
		if err != nil {
			panic(err)
		}
		msg := "\n>> A url " + url + " possui " + strconv.Itoa(len(links)) + " links."
		fmt.Println(msg)
		geraLog(msg)
		extraiLinks(doc)
	}
}

func extraiLinks(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key != "href" {
				continue
			}
			link, err := url.Parse(attr.Val)
			if err != nil || link.Scheme == "" {
				continue
			}
			links = append(links, link.String())
			msg := "\tVisitando >> " + link.String()
			fmt.Println(msg)
			geraLog(msg)
			verificaLink(link.String())
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extraiLinks(c)
	}
}

//geraLog salva em arquivo texto de log a atividade parâmetro 'mensagem'.
func geraLog(mensagem string) {
	arq, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao abrir arquivo de log - o mesmo não será registrado nesse ciclo de testes")
	} else {
		arq.WriteString(mensagem + "\n")
	}
	arq.Close()
}

func main() {
	url := menu()
	if len(url) == 0 {
		fmt.Println("Encerrando...")
		os.Exit(0)
	} else {
		verificaLink(url)
	}
}
