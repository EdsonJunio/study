package main

import "fmt"

func main() {
	text := "radar"
	isPalindrome := true

	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("É palíndromo")
	} else {
		fmt.Println("Não é palíndromo")
	}
}
