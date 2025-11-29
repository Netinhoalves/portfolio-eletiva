package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)

	if numero%2 == 0 {
		fmt.Println(numero, "é um número par")
	} else {
		fmt.Println(numero, "é número ímpar")
	}

	if numero > 0 {
		fmt.Println(numero, "é um número positivo")
	} else if numero < 0 {
		fmt.Println(numero, "é um número negativo")
	} else {
		fmt.Println("O número é zero")
	}
}