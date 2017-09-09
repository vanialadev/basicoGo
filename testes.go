package main

import (
	"fmt"
	//"reflect"
)

func main(){

	nome := devolveNome()
	fmt.Println(nome)

	nome, idade := devolveNomeEIdade()
	fmt.Println("Meu nome é", nome, "e tenho", idade, "anos")

	nome2, _ := devolveNomeEIdade()
	fmt.Println("Meu nome é", nome2)

	_, idade2 := devolveNomeEIdade()
	fmt.Println("Tenho", idade2, "anos")



	//var oi string
	//nome := ""
	//
	//fmt.Println("valor default da string é", oi)
	//if oi == ""{
	//	fmt.Println("olha só, o valor default de uma string em go é \"\"")
	//}
	//var i int
	//fmt.Println("valor default de int é", i)
	//var o bool
	//fmt.Println("valor default de bool é", o)
	//var p float64
	//fmt.Println("valor default de float64 p", p)
	//
	//
	//fmt.Println("A variável nome é do tipo ", reflect.TypeOf(nome))
	//fmt.Println("O endereço de memória da minha variável comando é", &o)
	//
	//var teste int
	//fmt.Scanf("%d", &teste)
	//fmt.Println("valor de memória de teste", &teste)
	//
	//var ola int
	//fmt.Scan(&ola)
	//fmt.Println("valor de memória de olá é ", &ola)
	//



}

func devolveNome() string{
	nome := "Vania"
	return nome
}

func devolveNomeEIdade() (string, int){
	nome := "Vania"
	idade := 27
	return nome, idade
}

func exibeNomes(){
	nomes := []string{"Vania", "Ana",  "Zé"}
	fmt.Println(nomes)
}