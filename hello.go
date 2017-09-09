package main

import (
	"fmt"
	"os"
	"net/http"
	"time"
	"bufio"
	"io"
	"strings"
	"strconv"
	"io/ioutil"
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
			imprimeLogs()
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
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "com probelmas. StatusCode é de", response.StatusCode)
		registraLog(site, false)
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

func registraLog(site string, status bool){

	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "= online " + strconv.FormatBool(status) + "\n")

	arquivo.Close()

}

func imprimeLogs(){

	arquivo, err := ioutil.ReadFile("logs.txt")

	fmt.Println(string(arquivo))

	if err != nil {
		fmt.Println(err)
	}

}