package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, _ := os.Getwd()
	caminho := filepath.Join(dir, "exemplo.ppm")

	f, err := os.Create(caminho)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	defer f.Close()

	conteudo := `P3
4 4
15
0  0  0    0  0  0    0  0  0   15  0 15
0  0  0    0 15  7    0  0  0    0  0  0
0  0  0    0  0  0    0 15  7    0  0  0
15  0 15    0  0  0    0  0  0    0  0  0
`
	fmt.Fprint(f, conteudo)

	fmt.Println("Arquivo 'exemplo.ppm' gerado com sucesso.")
}