package ocp

import "fmt"

type (
	Invoice struct {
		Value float64
	}

	CreditCard struct{}
	PayPal     struct{}
	Pix        struct{}
	Pross      struct{}

	PaymentMethod interface {
		Process(p *Invoice) error
	}
)

func (cc *CreditCard) Process(p *Invoice) error {
	fmt.Printf("Processing R$%.2f via credit card...\n", p.Value)
	return nil
}

func (pp *PayPal) Process(p *Invoice) error {
	fmt.Printf("Processing R$%.2f via PayPal...\n", p.Value)
	return nil
}

func (pix *Pix) Process(p *Invoice) error {
	fmt.Printf("Generating QR Code for payment of R$%.2f via Pix...\n", p.Value)
	return nil
}

func (po *Pross) Process(p *Invoice, pm PaymentMethod) error {
	err := pm.Process(p)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	proc := &Invoice{
		Value: 100.00,
	}

	pro := Pross{}

	err := pro.Process(proc, &CreditCard{})
	if err != nil {
		fmt.Println("Erro:", err)
	}
}
