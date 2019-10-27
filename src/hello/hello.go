package main

import (
	"fmt"
	"reflect"
)

func main() {

	//não é necessário incluir o tipo da variável na declaração
	var nome = "Fernando"
	var versao float32 = 1.1

	//declaração de variável sem o var
	idade := 22

	fmt.Println("Olá senhor ", nome, "sua idae é", idade)
	fmt.Println("Este programa está na versão", versao)

	//para descobrir o tipo da variável
	fmt.Println("O tipo da variável é", reflect.TypeOf(nome))

}
