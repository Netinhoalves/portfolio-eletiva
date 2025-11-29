package main

import "fmt"

func calcularAreaRetangulo(base, altura float64) float64 {
	return base * altura
}

func main() {
	var base, altura float64

	fmt.Print("Digite o valor da base do retângulo: ")
	fmt.Scanln(&base)

	fmt.Print("Digite o valor da altura do retângulo: ")
	fmt.Scanln(&altura)

	area := calcularAreaRetangulo(base, altura)

	fmt.Printf("A área de um retângulo com base %.2f e altura %.2f é %.2f\n", base, altura, area)
}