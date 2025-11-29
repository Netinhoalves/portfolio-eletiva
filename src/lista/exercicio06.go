package main

import (
	"fmt"
	"sort"
)

func main() {
	numeros := []int{9, 5, 2, 7, 1, 8, 3, 6, 4}

	fmt.Println("Sequência original:", numeros)

	sort.Ints(numeros)

	fmt.Println("Sequência ordenada:", numeros)
}