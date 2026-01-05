package main

import (
	"errors"
	"fmt"
)

func ReadFile(filename string) error {
	return errors.New("arquivo inexistente")
}

func LoadConfig() error {
	err := ReadFile("config.json")
	if err != nil {
		return fmt.Errorf("falha ao carregar configuração: %w", err)

	}
	return nil
}

func main() {
	if err := LoadConfig(); err != nil {
		fmt.Println(err)
	}
}
