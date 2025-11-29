package main
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("--- 26. Algoritmo Genético (Simplificado) ---")
	alvo := "GOLANG"
	genes := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	atual := "AAAAAA" 
	geracao := 0
	
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Meta: %s\n", alvo)
	
	for atual != alvo {
		geracao++
		// Mutação
		r := []rune(atual)
		idx := rand.Intn(len(alvo))
		if r[idx] != rune(alvo[idx]) {
			r[idx] = rune(genes[rand.Intn(len(genes))])
		}
		atual = string(r)
		
		if geracao%10 == 0 {
			fmt.Printf("Gen %d: %s\n", geracao, atual)
		}
	}
	fmt.Printf("FINAL: %s atingido na geração %d\n", atual, geracao)
}