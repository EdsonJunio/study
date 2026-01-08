package main

import (
	"errors"
	"fmt"
)

func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("não pode dividir por zero")
	}
	return a / b, nil
}
func main() {
	soma, err := dividir(10, 20)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(soma)
}
