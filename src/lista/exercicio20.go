package main

import (
	"fmt"
	"strings"
)

func main() {
	texto := "Go é uma linguagem de programação. A linguagem Go foi criada pelo Google. Aprender Go é divertido."
	palavra := "Go"

	ocorrencias := strings.Count(texto, palavra)

	fmt.Printf("O texto é: '%s'\n", texto)
	fmt.Printf("A palavra '%s' aparece %d vezes no texto.\n", palavra, ocorrencias)

	textoLower := strings.ToLower(texto)
	palavraLower := strings.ToLower(palavra)
	ocorrenciasInsensitive := strings.Count(textoLower, palavraLower)
	fmt.Printf("(Busca case-insensitive) A palavra '%s' aparece %d vezes.\n", palavra, ocorrenciasInsensitive)
}