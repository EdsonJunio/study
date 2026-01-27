package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i <= n/i; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("5 é primo?", isPrime(5))
	fmt.Println("10 é primo?", isPrime(10))
}
