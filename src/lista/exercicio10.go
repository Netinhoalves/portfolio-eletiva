package main

import "fmt"

func hanoi(n int, origem, destino, auxiliar string) {
	if n == 1 {
		fmt.Printf("Mova o disco 1 da torre %s para a torre %s\n", origem, destino)
		return
	}

	hanoi(n-1, origem, auxiliar, destino)

	fmt.Printf("Mova o disco %d da torre %s para a torre %s\n", n, origem, destino)

	hanoi(n-1, auxiliar, destino, origem)
}

func main() {
	numeroDeDiscos := 3
	fmt.Printf("Resolvendo Torres de Hanói para %d discos:\n", numeroDeDiscos)
	hanoi(numeroDeDiscos, "A", "C", "B")
}