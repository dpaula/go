package main

import (
	"fmt"
	"os"
)

func main() {

	exibeInformacoes()

	exibeMenu()

	comando := leComando()

	// no go, o switch não precisa de break
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 3:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Comando não existe")
		os.Exit(-1)
	}
}

func exibeInformacoes() {

	//não é necessário incluir o tipo da variável na declaração
	var nome = "Fernando"
	var versao float32 = 1.1

	fmt.Println("Olá senhor ", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair")
}

func leComando() int {

	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido

}
