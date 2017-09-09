package main

import (
	"fmt"
	"os"
	"net/http"
	"time"
	"bufio"
	"io"
	"strings"
)

const monitoramento = 3
const delay = 5

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Não reconheço esse comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao(){
	nome := "Vânia"
	versao := 1.1
	fmt.Println("Olá, sra.", nome)
	fmt.Println("A versão do programa é ", versao)
}

func exibeMenu(){
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int{
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("Você escolheu a opção", comandoLido)
	fmt.Println("")
	return comandoLido
}

func iniciaMonitoramento(){
	fmt.Println("Monitorando...")

	sites := leSitesDoArquivo()

	for i:= 0; i < monitoramento ; i++  {
		for i, site := range sites{
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func testaSite(site string){
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	fmt.Println(response)
	if response.StatusCode == 200 {
		fmt.Println("Site", site, "carregado com sucesso")
	} else {
		fmt.Println("Site", site, "com probelmas. StatusCode é de", response.StatusCode)
	}
}

func leSitesDoArquivo() []string{

	var sites[] string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
	return sites
}