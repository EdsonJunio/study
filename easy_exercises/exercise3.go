package main

import (
	"errors"
	"fmt"
)

func FindMax(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("the list cannot be")
	}

	maxVal := numbers[0]

	for _, num := range numbers {
		if num > maxVal {
			maxVal = num
		}
	}

	return maxVal, nil
}

func main() {
	numbers := []int{1, 2, 3, 10, 5, 6, 7, 8, 9}

	maxVal, err := FindMax(numbers)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("The highest value is:", maxVal)
	}
}
