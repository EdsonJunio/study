package main

import "fmt"

func Sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Enter two numbers:")
	var a, b int

	fmt.Scan(&a, &b)
	result := Sum(a, b)

	fmt.Printf("The sum is %d\n", result)

}
