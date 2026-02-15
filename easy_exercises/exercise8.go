package easy_exercises

import "fmt"

func main() {
	var number int
	sum := 1
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		sum *= i
	}
	fmt.Println(sum)
}
