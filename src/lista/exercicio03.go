package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)

	antecessor := numero - 1
	sucessor := numero + 1

	fmt.Println("O antecessor de", numero, "é:", antecessor)
	fmt.Println("O sucessor de", numero, "é:", sucessor)
}