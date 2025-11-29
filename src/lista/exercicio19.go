package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	texto := "Análise e Desenvolvimento de Sistemas"
	vogais := 0
	consoantes := 0

	texto = strings.ToLower(texto)

	for _, char := range texto {
		if unicode.IsLetter(char) {
			switch char {
			case 'a', 'e', 'i', 'o', 'u', 'á', 'é', 'í', 'ó', 'ú', 'à', 'ã', 'õ', 'ê', 'ô':
				vogais++
			default:
				consoantes++
			}
		}
	}

	fmt.Printf("No texto '%s', temos:\n", texto)
	fmt.Println("Vogais:", vogais)
	fmt.Println("Consoantes:", consoantes)
}