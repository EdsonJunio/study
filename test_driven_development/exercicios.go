package main

func Subtrai(a, b int) int {
	return a - b
}

func Multiplica(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	return a / b
}

func Dobro(n int) int {
	return n * 2
}

func Ola(nome string) string {
	return "Olá " + nome
}

func ContaCaracteres(texto string) int {
	return len(texto)
}

func EPar(num int) bool {
	return num%2 == 0
}

func EhMaiorDeIdade(idade int) bool {
	return idade >= 18
}

func ConverteCelsiusParaFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func PerimetroRetangulo(base, altura float64) float64 {
	return 2 * (base + altura)
}
