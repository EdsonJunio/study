package easy_exercises

import "fmt"

func main() {
	var valueA, soma int
	fmt.Scan(&valueA)

	for valueA > 0 {
		digito := valueA % 10
		soma += digito
		valueA /= 10
	}
	fmt.Println(soma)
}
