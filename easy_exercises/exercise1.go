package main

import (
	"fmt"
)

func CalculateSum(limit int) int {
	sum := 0
	for i := 1; i <= limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	limit := 10
	result := CalculateSum(limit)
	fmt.Printf("A soma dos múltiplos de 3 ou 5 até %d é: %d\n", limit, result)
}
