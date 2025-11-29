package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	numeroSecreto := rand.Intn(100) + 1
	var chute int
	tentativas := 0

	fmt.Println("--- Jogo da Adivinhação ---")
	fmt.Println("Eu pensei em um número entre 1 e 100. Tente adivinhar!")

	for {
		fmt.Print("Qual é o seu chute? ")
		fmt.Scanln(&chute)
		tentativas++

		if chute < numeroSecreto {
			fmt.Println("Muito baixo! Tente um número maior.")
		} else if chute > numeroSecreto {
			fmt.Println("Muito alto! Tente um número menor.")
		} else {
			fmt.Printf("Parabéns! Você acertou o número %d em %d tentativas!\n", numeroSecreto, tentativas)
			break
		}
	}
}