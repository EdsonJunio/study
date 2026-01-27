package main

import "fmt"

func CalculateOperation(value int64) (int64, int64) {
	multiplied := value * 2
	tax := int64(100)

	return multiplied, tax
}

func main() {
	profit, tax := CalculateOperation(5000)

	fmt.Printf("Lucro: %d, Taxa: %d\n", profit, tax)
}
