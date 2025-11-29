package main

import "fmt"

func calcularIMC(pesoKg, alturaM float64) float64 {
	if alturaM <= 0 {
		return 0
	}
	return pesoKg / (alturaM * alturaM)
}

func classificarIMC(imc float64) string {
	if imc < 18.5 {
		return "Abaixo do peso"
	} else if imc < 24.9 {
		return "Peso normal"
	} else if imc < 29.9 {
		return "Sobrepeso"
	} else if imc < 34.9 {
		return "Obesidade grau 1"
	} else if imc < 39.9 {
		return "Obesidade grau 2"
	} else {
		return "Obesidade grau 3 (mórbida)"
	}
}

func main() {
	var peso, altura float64

	fmt.Print("Digite seu peso em kg (ex: 70.5): ")
	fmt.Scanln(&peso)
	fmt.Print("Digite sua altura em metros (ex: 1.75): ")
	fmt.Scanln(&altura)

	imc := calcularIMC(peso, altura)
	classificacao := classificarIMC(imc)

	fmt.Printf("Seu IMC é: %.2f\n", imc)
	fmt.Printf("Classificação: %s\n", classificacao)
}