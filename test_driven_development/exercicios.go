package test_driven_development

func Subtrai(a, b int) int {
	return a + b
}

func Multiplica(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	return a / b
}

func Dobro(a int) int {
	return a * 2
}

func Ola(nome string) string {
	return nome
}

func TamanhoDaString(text string) int {
	tamanho := 0

	for i := 0; i < len(text); i++ {
		tamanho++
	}

	return tamanho
}

func Par(a int) bool {

	if a%2 == 0 {
		return true
	}

	return false
}
