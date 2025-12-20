package main

import (
	"errors"
	"fmt"
)

type Processor interface {
	Authorize()
}

type Debit struct {
	AccountBalance float64
}

type Credit struct{}

func (d *Debit) Authorize() {
	fmt.Printf("Debit authorized, Remaining balance $%.2f\n", d.AccountBalance)
}

func (Credit) Authorize() {
	fmt.Printf("Credit authorized")
}

func NewProcessor(typ string, currentBalance float64) (Processor, error) {

	if typ == "credit" {
		return Credit{}, nil
	} else if typ == "debit" && currentBalance < 100 {
		return nil, errors.New("insufficient funds to start operation")
	} else if typ == "debit" {
		return &Debit{AccountBalance: currentBalance}, nil
	}

	return nil, errors.New("invalid processor type")
}

func main() {
	pay, err := NewProcessor("credit", 100)
	if err != nil {
		fmt.Println("error", err)
	}

	pay.Authorize()
}
