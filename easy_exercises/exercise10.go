package easy_exercises

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	if number <= 1 {
		fmt.Println("Não Primo")
		return
	}

	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			fmt.Println("Não Primo")
			return
		}
	}

	fmt.Println("Primo")
}
