package easy_exercises

import "fmt"

func main() {
	var N, T, soma int
	fmt.Scan(&N)

	sequencia := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&sequencia[i])
	}

	fmt.Scan(&T)

	for i := 0; i < N; i++ {
		if sequencia[i] > T {
			soma += sequencia[i]
		}
	}

	fmt.Println(soma)
}
