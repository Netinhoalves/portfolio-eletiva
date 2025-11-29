package main

import (
	"fmt"
	"math"
)

func main() {
	var numero int
	fmt.Print("Digite um número inteiro para verificar se é primo: ")
	fmt.Scanln(&numero)

	primo := true
	
	if numero <= 1 {
		primo = false
	} else {
		limite := int(math.Sqrt(float64(numero)))
		for i := 2; i <= limite; i++ {
			if numero%i == 0 {
				primo = false
				break
			}
		}
	}

	if primo {
		fmt.Println(numero, "é um número primo.")
	} else {
		fmt.Println(numero, "não é um número primo.")
	}
}