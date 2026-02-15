package easy_exercises

import "fmt"

func main() {
	var A, B int

	fmt.Scan(&A, &B)

	soma := A + B

	if A == B {
		fmt.Println(soma * 2)
	} else {
		fmt.Println(soma)
	}
}
