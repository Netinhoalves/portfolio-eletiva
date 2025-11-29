package main

import "fmt"

func celsiusParaFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitParaCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	var temp float64
	var unidade string

	fmt.Print("Digite a temperatura (ex: 25): ")
	fmt.Scanln(&temp)
	fmt.Print("Digite a unidade de origem ('C' para Celsius ou 'F' para Fahrenheit): ")
	fmt.Scanln(&unidade)

	switch unidade {
	case "C", "c":
		fahrenheit := celsiusParaFahrenheit(temp)
		fmt.Printf("%.2f°C equivale a %.2f°F\n", temp, fahrenheit)
	case "F", "f":
		celsius := fahrenheitParaCelsius(temp)
		fmt.Printf("%.2f°F equivale a %.2f°C\n", temp, celsius)
	default:
		fmt.Println("Unidade inválida. Por favor, use 'C' ou 'F'.")
	}
}