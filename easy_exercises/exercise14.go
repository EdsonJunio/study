package easy_exercises

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var numero int
	var maior int

	for i := 0; i < n; i++ {
		fmt.Scan(&numero)

		if i == 0 || numero > maior {
			maior = numero
		}
	}

	fmt.Println(maior)
}
