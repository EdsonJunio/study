package main

import "testing"

func TestSubtrai(t *testing.T) {
	// arrange
	numA := 10
	numB := 5
	result := 5

	// act
	subtraiObtida := Subtrai(numA, numB)

	// assert
	if subtraiObtida != result {
		t.Errorf("Esperado '%d', mas obteve '%d'\n", result, subtraiObtida)
	}
}

func TestMultiplicador(t *testing.T) {
	valueA := 7
	valueB := 7
	resultMultiplos := 49

	multiplicadorObtido := Multiplica(valueA, valueB)

	if multiplicadorObtido != resultMultiplos {
		t.Errorf("valor esperado '%d' valor recebudo '%d'\n", resultMultiplos, multiplicadorObtido)
	}
}

func TestDivisao(t *testing.T) {
	valueA := 100
	valueB := 10
	resultDivisao := 10

	divisaoObtida := Divide(valueA, valueB)

	if divisaoObtida != resultDivisao {
		t.Errorf("Valor esperado '%d' valor recebido '%d'\n", resultDivisao, divisaoObtida)
	}

}

func TestDobro(t *testing.T) {
	valueA := 10
	respostaDobro := 20

	dobroObtido := Dobro(valueA)

	if dobroObtido != respostaDobro {
		t.Errorf("Valor esperado '%d' valor obtido '%d'", respostaDobro, dobroObtido)
	}
}

func TestNome(t *testing.T) {
	valueA := "Edson"
	resultadoEsperado := "Olá Edson"

	nomeRecebido := Ola(valueA)

	if nomeRecebido != resultadoEsperado {
		t.Errorf("Valor esperado '%s' valor obtido '%s'", resultadoEsperado, nomeRecebido)
	}
}

func TestContaCaracteres(t *testing.T) {
	nome := "Edson"
	esperado := 5

	obtido := ContaCaracteres(nome)

	if obtido != esperado {
		t.Errorf("Esperado '%d', obtido '%d'", esperado, obtido)
	}
}

func TestEpar(t *testing.T) {
	valueA := 4
	valueEsperado := true

	valueObtido := EPar(valueA)

	if valueObtido != valueEsperado {
		t.Errorf("Valor esperado '%v' valor obtido '%v'", valueEsperado, valueObtido)
	}
}

func TestEhMaiorDeIdade(t *testing.T) {
	idade := 18
	esperado := true

	obtido := EhMaiorDeIdade(idade)

	if obtido != esperado {
		t.Errorf("Esperado '%v', obtido '%v' para idade %d", esperado, obtido, idade)
	}
}

func TestConverteCelsiusParaFahrenheit(t *testing.T) {
	celsius := 100.0
	esperado := 212.0

	obtido := ConverteCelsiusParaFahrenheit(celsius)

	if obtido != esperado {
		t.Errorf("Esperado '%.2f', obtido '%.2f'", esperado, obtido)
	}
}

func TestPerimetroRetangulo(t *testing.T) {
	valueBase := 10.0
	valueAltura := 5.0
	ValueEsperado := 30.0

	valueRecebido := PerimetroRetangulo(valueBase, valueAltura)

	if valueRecebido != ValueEsperado {
		t.Errorf("Valor esperado '%f' valor obtido '%f'", ValueEsperado, valueRecebido)
	}
}
