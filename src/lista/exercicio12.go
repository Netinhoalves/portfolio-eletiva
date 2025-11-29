package main

import "fmt"

func saoIguais(a, b int) bool {
	return a == b
}

func main() {
	valor1 := 10
	valor2 := 10
	valor3 := 20

	fmt.Printf("Os valores %d e %d são iguais? %t\n", valor1, valor2, saoIguais(valor1, valor2))
	fmt.Printf("Os valores %d e %d são iguais? %t\n", valor1, valor3, saoIguais(valor1, valor3))

	string1 := "go"
	string2 := "golang"
	fmt.Printf("As strings \"%s\" e \"%s\" são iguais? %t\n", string1, string2, string1 == string2)
}