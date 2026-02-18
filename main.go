package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	numeros := make([]int, n)
	fmt.Println(numeros)

	for i := 0; i < n; i++ {
		fmt.Scan(&numeros[i])
	}

	for i := n - 1; i >= 0; i-- {
		fmt.Print(numeros[i], " ")
	}
}
