package main

func ProcessPayment(active bool, balance float64, amount float64) string {

	if !active {
		return "Erro: Conta inativa"
	}

	if amount <= 0 {
		return "Erro: Valor inválido"
	}

	if balance < amount {
		return "Erro: Saldo insuficiente"
	}

	return "Pagamento processado!"
}
