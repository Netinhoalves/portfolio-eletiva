package main

import "fmt"

func main() {
	nome := "Olá, mundo!"

	fmt.Println("O valor da variável 'nome' é:", nome)

	fmt.Println("O endereço de memória da variável 'nome' é:", &nome)
}