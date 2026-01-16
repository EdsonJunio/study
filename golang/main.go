package main

import "fmt"

func main() {
	var saldo int64 = 100000
	var boleto int64 = 100000

	if novoSaldo := saldo - boleto; novoSaldo >= 0 {

		fmt.Println("Pago. Restante:", novoSaldo)
	} else {
		fmt.Println("Falha")
	}
}
