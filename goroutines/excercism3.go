package goroutines

import "fmt"

func main() {
	ch := make(chan string, 3)

	ch <- "Pagamento 1"
	ch <- "Pagamento 2"
	ch <- "Pagamento 3"

	fmt.Println("Enviados para o buffer sem travar!")

	fmt.Println(<-ch)

	ch <- "Pagamento 4"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
