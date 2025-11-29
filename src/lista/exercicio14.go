package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ehPalindromo(texto string) bool {
	var builder strings.Builder
	for _, char := range texto {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			builder.WriteRune(unicode.ToLower(char))
		}
	}
	textoLimpo := builder.String()

	runas := []rune(textoLimpo)
	n := len(runas)

	for i := 0; i < n/2; i++ {
		if runas[i] != runas[n-1-i] {
			return false
		}
	}

	return true
}

func main() {
	frase1 := "A sacada da casa"
	frase2 := "ovo"
	frase3 := "Go Language"

	fmt.Printf("'%s' é um palíndromo? %t\n", frase1, ehPalindromo(frase1))
	fmt.Printf("'%s' é um palíndromo? %t\n", frase2, ehPalindromo(frase2))
	fmt.Printf("'%s' é um palíndromo? %t\n", frase3, ehPalindromo(frase3))
}