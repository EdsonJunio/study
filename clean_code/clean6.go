package main

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance float64
}

func NewAccount(owner string, initialBalance float64) (*Account, error) {
	if owner == "" {
		return nil, errors.New("owner não pode ser vazio")
	}

	if initialBalance < 0 {
		return nil, errors.New("saldo inicial não pode ser negativo")
	}

	return &Account{
		owner:   owner,
		balance: initialBalance,
	}, nil
}

func main() {
	acc, err := NewAccount("Edson", 100)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Conta criada com sucesso:", acc)
}
