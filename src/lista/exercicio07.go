package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	texto := "algoritmo"
	fmt.Println("Texto original:", texto)

	caracteres := strings.Split(texto, "")

	sort.Strings(caracteres)

	textoOrdenado := strings.Join(caracteres, "")

	fmt.Println("Texto com caracteres ordenados:", textoOrdenado)
}