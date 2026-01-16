package main

import "testing"

func TestContaDeposita(t *testing.T) {
	t.Run("Faz um deposito", func(t *testing.T) {
		// Arrange
		conta := Conta{Saldo: 500}
		valorDeposito := 200
		resultadoEsperado := 700.0

		// Act
		conta.Deposita(float64(valorDeposito))
		obtido := conta.VerificaSaldo()

		// Assert
		if conta.Saldo != resultadoEsperado {
			t.Errorf(
				"valor esperado %.2f, valor recebido %.2f",
				resultadoEsperado,
				conta.Saldo,
			)
		}

		if obtido != resultadoEsperado {
			t.Errorf(
				"valor esperado %.2f, valor recebido %.2f",
				resultadoEsperado,
				conta.VerificaSaldo(),
			)
		}
	})
}
