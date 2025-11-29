package main

import "fmt"

func main() {
	var numero1 float64 = 25.0
	var numero2 float64 = 4.0

	if numero2 == 0 {
		fmt.Println("Erro: Divisão por zero não é permitida.")
	} else {
		divisao := numero1 / numero2
		fmt.Println("O resultado da divisão de", numero1, "por", numero2, "é", divisao)
	}
}