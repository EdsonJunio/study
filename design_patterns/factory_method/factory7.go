package main

import "fmt"

type Gateway interface {
	Charge(value float64)
}

type Stripe struct{}

type FakeGateway struct{}

func (Stripe) Charge(value float64) {
	fmt.Printf("Cobrando $ %.2f na API do Stripe", value)
}

func (FakeGateway) Charge(value float64) {
	fmt.Printf("Fake: Registrando cobrança de %.2f apenas no log", value)
}

func NewGateway(isProduction bool) Gateway {

	switch isProduction {
	case true:
		return &Stripe{}
	case false:
		return &FakeGateway{}
	}

	return nil
}

func main() {
	gateway := NewGateway(false)
	gateway.Charge(100)
}
