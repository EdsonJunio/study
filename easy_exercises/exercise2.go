package easy_exercises

import "fmt"

func main() {
	var A int
	var N = 21
	fmt.Scan(&A)

	if A <= N {
		var sub = N - A
		fmt.Println(sub)
	} else {
		var total = A - N
		fmt.Println(total * 2)
	}

}
