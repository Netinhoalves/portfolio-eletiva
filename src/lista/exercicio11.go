package main

import (
	"fmt"
	"time"
)

func main() {
	var dataNascimento string
	fmt.Print("Digite sua data de nascimento (formato DD/MM/AAAA): ")
	fmt.Scanln(&dataNascimento)

	const layout = "02/01/2006"
	data, err := time.Parse(layout, dataNascimento)

	if err != nil {
		fmt.Println("Formato de data inválido. Use DD/MM/AAAA. Erro:", err)
		return
	}

	diaDaSemana := data.Weekday()

	diasEmPortugues := map[time.Weekday]string{
		time.Sunday:    "Domingo",
		time.Monday:    "Segunda-feira",
		time.Tuesday:   "Terça-feira",
		time.Wednesday: "Quarta-feira",
		time.Thursday:  "Quinta-feira",
		time.Friday:    "Sexta-feira",
		time.Saturday:  "Sábado",
	}

	fmt.Printf("Você nasceu em um(a) %s.\n", diasEmPortugues[diaDaSemana])
}