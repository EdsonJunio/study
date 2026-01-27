package main

import "fmt"

func CalculateDebt(debt int64, heritage int64) {

	if percentage := (debt * 100) / heritage; percentage > 50 {
		fmt.Printf("Alavancagem Alta:[%d]%%", percentage)
	} else {
		fmt.Printf("Alavancagem Segura:[%d]%%\n", percentage)
	}

}

func main() {
	var debt int64 = 60000
	var heritage int64 = 100000
	CalculateDebt(debt, heritage)
}
