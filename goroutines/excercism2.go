package goroutines

import "fmt"

func main() {
	canal := make(chan float64)
	go calcularRendimento(5000.0, canal)
	fmt.Println("Aguardando processamento...")
	resultado := <-canal
	fmt.Println(resultado)
}

func calcularRendimento(valorBase float64, ch chan float64) {
	rendimento := valorBase * 1.10
	ch <- rendimento
}
