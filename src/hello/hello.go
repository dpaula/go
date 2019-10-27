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
	fmt.Scanf("%d", &comando)

	//para mostrar o endereço da variável
	fmt.Println("Endereço da variável", &comando)

	fmt.Println("Comando escolhido", comando)

}
