package srp

import (
	"fmt"
	"strconv"
	"strings"
)

type Bot struct{}

// HandleMessage decide E executa. Violação de SRP.
func (b *Bot) HandleMessage(msg string) {
	parts := strings.Split(msg, " ")
	command := parts[0]

	switch command {
	case "/hello":
		// Lógica de saudação
		fmt.Println("Olá! Sou um bot SRP.")

	case "/soma":
		// Lógica de cálculo + Validação + Print
		if len(parts) < 3 {
			fmt.Println("Erro: use /soma 10 20")
			return
		}
		v1, _ := strconv.Atoi(parts[1])
		v2, _ := strconv.Atoi(parts[2])
		res := v1 + v2
		fmt.Printf("Resultado: %d\n", res)

	case "/reverse":
		// Lógica de manipulação de string
		if len(parts) < 2 {
			return
		}
		word := parts[1]
		rns := []rune(word)
		for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
			rns[i], rns[j] = rns[j], rns[i]
		}
		fmt.Println(string(rns))

	default:
		fmt.Println("Comando desconhecido")
	}
}
