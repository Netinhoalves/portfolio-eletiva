package main

import "fmt"

func fatorialIterativo(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	var resultado uint64 = 1
	var i uint64
	for i = 2; i <= n; i++ {
		resultado *= i
	}
	return resultado
}

func fatorialRecursivo(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * fatorialRecursivo(n-1)
}

func main() {
	var numero uint64 = 20

	fmt.Printf("O fatorial de %d (iterativo) é: %d\n", numero, fatorialIterativo(numero))
	fmt.Printf("O fatorial de %d (recursivo) é: %d\n", numero, fatorialRecursivo(numero))
}