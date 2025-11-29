package main

import "fmt"

func calcularMedia(numeros ...float64) float64 {
	if len(numeros) == 0 {
		return 0.0
	}

	total := 0.0
	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func main() {
	media1 := calcularMedia(10, 20)
	fmt.Printf("A média de 10 e 20 é: %.2f\n", media1)

	media2 := calcularMedia(5.5, 10.0, 8.5, 6.0)
	fmt.Printf("A média de 5.5, 10.0, 8.5 e 6.0 é: %.2f\n", media2)

	outrosNumeros := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	media3 := calcularMedia(outrosNumeros...)
	fmt.Printf("A média da slice é: %.2f\n", media3)
}