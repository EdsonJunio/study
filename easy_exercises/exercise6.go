package easy_exercises

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	for i := 0; i <= number; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
