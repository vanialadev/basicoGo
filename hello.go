package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
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
	return comandoLido
}
