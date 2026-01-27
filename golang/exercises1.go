package main

import "fmt"

const TaxaCambio float64 = 5.25

func main() {
	var valorDolar int64 = 10000
	mult := int64(float64(valorDolar) * TaxaCambio)
	fmt.Printf("Valor em Reais (centavos) ", mult)
}
