package main

import "fmt"

func main() {

	// retorno de duas variáveis
	nome, idade := retornaNome()

	fmt.Println("Seu nome é", nome, "idade", idade)

	// retorno da função ignorando uma variável
	_, diaMes := retornaDiaSemanaEMes()

	fmt.Println("O dia do mes é", diaMes)

}

// Função que retonra duas variáveis
func retornaNome() (string, int) {

	nome := "Fernando"
	idade := 34

	return nome, idade
}

// Função que retonra duas variáveis
func retornaDiaSemanaEMes() (string, int) {

	diaSemana := "Domingo"
	diaMes := 10

	return diaSemana, diaMes
}
