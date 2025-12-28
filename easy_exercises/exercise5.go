package main

import (
	"fmt"
)

func CountVowels(text string) int {
	count := 0

	for _, char := range text {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}

	return count
}

func main() {
	input := "Golang"
	total := CountVowels(input)
	fmt.Printf("A palavra '%s' tem %d vogais.\n", input, total)
}
