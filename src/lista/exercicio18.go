package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func main() {
	const largura, altura = 256, 256

	matrizR := make([][]uint8, altura)
	matrizG := make([][]uint8, altura)
	matrizB := make([][]uint8, altura)

	for y := 0; y < altura; y++ {
		matrizR[y] = make([]uint8, largura)
		matrizG[y] = make([]uint8, largura)
		matrizB[y] = make([]uint8, largura)
		for x := 0; x < largura; x++ {
			matrizR[y][x] = uint8(x)
			matrizG[y][x] = uint8(y)
			matrizB[y][x] = 128
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, largura, altura))

	for y := 0; y < altura; y++ {
		for x := 0; x < largura; x++ {
			r := matrizR[y][x]
			g := matrizG[y][x]
			b := matrizB[y][x]

			pixel := color.RGBA{R: r, G: g, B: b, A: 255}

			img.Set(x, y, pixel)
		}
	}

	nomeArquivo := "imagem_gerada.jpg"
	arquivo, err := os.Create(nomeArquivo)
	if err != nil {
		panic(err)
	}
	defer arquivo.Close()

	opcoes := jpeg.Options{Quality: 90}
	err = jpeg.Encode(arquivo, img, &opcoes)
	if err != nil {
		panic(err)
	}

	fmt.Println("Arquivo", nomeArquivo, "criado com sucesso!")
}