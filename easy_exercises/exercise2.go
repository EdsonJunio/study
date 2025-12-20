package main

import (
	"fmt"
)

func CalculateFactorial(n int) (uint64, error) {
	if n < 0 {
		return 0, fmt.Errorf("fatorial não definido para números negativos: %d", n)
	}

	result := uint64(1)
	for i := 2; i <= n; i++ {
		result *= uint64(i)
	}

	return result, nil
}

func main() {
	n := 5
	fatorial, err := CalculateFactorial(n)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d! = %d\n", n, fatorial)
	}
}
