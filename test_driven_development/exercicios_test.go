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
