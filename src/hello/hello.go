package main

import (
	"fmt"
	"reflect"
)

func main() {

	var nome = "Fernando"
	var versao float32 = 1.1
	idade := 22

	fmt.Println("Olá senhor ", nome, "sua idae é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variável é", reflect.TypeOf(nome))

}
