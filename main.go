package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	numeros := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numeros[i])
	}

	crescente := true

	for i := 1; i < n; i++ {
		if numeros[i] <= numeros[i-1] {
			crescente = false
			break
		}
	}

	if crescente {
		fmt.Println("Crescente")
	} else {
		fmt.Println("Desordenado")
	}
}
