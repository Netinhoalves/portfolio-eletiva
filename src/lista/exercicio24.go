package main

import "fmt"

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func mmc(a, b int) int {
	return (a / mdc(a, b)) * b
}

func main() {
	num1 := 12
	num2 := 18

	resultadoMDC := mdc(num1, num2)
	resultadoMMC := mmc(num1, num2)

	fmt.Printf("Dados os números %d e %d:\n", num1, num2)
	fmt.Printf("O Máximo Divisor Comum (MDC) é: %d\n", resultadoMDC)
	fmt.Printf("O Mínimo Múltiplo Comum (MMC) é: %d\n", resultadoMMC)
}