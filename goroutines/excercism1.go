package goroutines

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		processar("Boleto")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		processar("TED")
	}()

	wg.Wait()
	fmt.Println("Fim do expediente")
}

func processar(item string) {

	for i := 0; i < 5; i++ {
		fmt.Println(item, i)
	}
}
