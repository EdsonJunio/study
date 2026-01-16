package main

import "fmt"

type Conta struct {
	Saldo float64
}

func (c *Conta) Deposita(valor float64) { // Mudou de ContaDeposita para Deposita
	c.Saldo += valor
}

func (c *Conta) VerificaSaldo() float64 {
	return c.Saldo
}

func main() {
	conta := Conta{}

	conta.Deposita(100)
	conta.Deposita(50)

	fmt.Println("Saldo:", conta.VerificaSaldo())
}
