package easy_exercises

import "fmt"

func main() {

	var valueA, valueB int
	fmt.Scan(&valueA, &valueB)
	sum := valueA + valueB

	if valueA == 10 || valueB == 10 || sum == 10 {
		fmt.Println("Verdadeiro")
		return
	}

	fmt.Println("Falso")
}
