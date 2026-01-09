package main

import "testing"

func TestSoma(t *testing.T) {
	// 1. Arrange (Preparo)
	numeroA := 5
	numeroB := 5
	resultadoEsperado := 10

	// 2. Act (Ação - Chamamos a função que queremos testar)
	// Nota: A função Soma ainda não existe, isso vai dar erro de compilação (que conta como RED)
	resultadoObtido := Soma(numeroA, numeroB)

	// 3. Assert (Verificação)
	if resultadoObtido != resultadoEsperado {
		// %d é placeholder para inteiros. TDD exige mensagens de erro úteis!
		t.Errorf("Esperado '%d', mas obteve '%d'", resultadoEsperado, resultadoObtido)
	}
}

func TestCalculaTotalVenda(t *testing.T) {
	qtd := 10
	precoUnitario := 5.0 // float
	esperado := 50.0

	obtido := CalculaTotal(qtd, precoUnitario)

	if obtido != esperado {
		t.Errorf("Erro no cálculo. Esperado %.2f, obtido %.2f", esperado, obtido)
	}
}
