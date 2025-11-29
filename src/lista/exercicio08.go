package main

import "fmt"

func main() {
	var escolha string

	fmt.Println("Você acorda em uma floresta. Há dois caminhos.")
	fmt.Print("Você vai pela 'esquerda' ou 'direita'? ")
	fmt.Scanln(&escolha)

	if escolha == "esquerda" {
		fmt.Println("\nVocê segue pela esquerda e encontra um rio.")
		fmt.Print("Você tenta 'atravessar' ou 'seguir' o rio? ")
		fmt.Scanln(&escolha)

		if escolha == "atravessar" {
			fmt.Println("\nA correnteza estava forte. Fim de jogo.")
		} else if escolha == "seguir" {
			fmt.Println("\nVocê segue o rio e encontra uma ponte segura. Você chegou a uma vila! Você venceu!")
		} else {
			fmt.Println("\nEscolha inválida. Você fica perdido. Fim de jogo.")
		}

	} else if escolha == "direita" {
		fmt.Println("\nVocê segue pela direita e encontra uma caverna escura.")
		fmt.Print("Você decide 'entrar' ou 'ignorar' a caverna? ")
		fmt.Scanln(&escolha)

		if escolha == "entrar" {
			fmt.Println("\nUm urso dormia na caverna. Ele não gostou de ser acordado. Fim de jogo.")
		} else if escolha == "ignorar" {
			fmt.Println("\nVocê ignora a caverna e encontra um campo aberto. Você está seguro! Você venceu!")
		} else {
			fmt.Println("\nEscolha inválida. Você fica perdido. Fim de jogo.")
		}

	} else {
		fmt.Println("\nVocê não escolheu um caminho e ficou parado. A noite chegou. Fim de jogo.")
	}
}