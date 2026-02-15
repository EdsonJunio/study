package easy_exercises

import "fmt"

func main() {
	var num, sum int
	fmt.Scan(&num)

	for i := 0; i <= num; i++ {
		sum += i
	}

	fmt.Println(sum)
}
