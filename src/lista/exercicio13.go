package main

import "fmt"

func encontrarModa(numeros []int) int {
	frequencias := make(map[int]int)

	for _, num := range numeros {
		frequencias[num]++
	}

	moda := -1
	maxFrequencia := 0

	for num, freq := range frequencias {
		if freq > maxFrequencia {
			maxFrequencia = freq
			moda = num
		}
	}

	return moda
}

func main() {
	sequencia := []int{1, 2, 3, 4, 2, 2, 5, 6, 2, 7, 8, 5}
	fmt.Println("Sequência:", sequencia)
	moda := encontrarModa(sequencia)
	fmt.Println("A moda da sequência é:", moda)
}