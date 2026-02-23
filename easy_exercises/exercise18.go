package easy_exercises

import "fmt"

func main() {
	var N, X int

	fmt.Scan(&N)

	numeros := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&numeros[i])
	}

	fmt.Scan(&X)

	posicao := -1

	for i := 0; i < N; i++ {
		if numeros[i] == X {
			posicao = i
			break
		}
	}

	fmt.Println(posicao)
}
