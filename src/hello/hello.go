package main

import (
	"fmt"
)

func main() {

	//não é necessário incluir o tipo da variável na declaração
	var nome = "Fernando"
	var versao float32 = 1.1

	fmt.Println("Olá senhor ", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair")

	var comando int

	//pegando valor no terminal e jogando na variável comando, usando & para apontar pra referencia em memória
	fmt.Scan(&comando)

	// no go, o switch não precisa de break
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 3:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Comando não existe")

	}

}
