package easy_exercises

import "fmt"

func main() {
	var valueA int
	var invertido = 0
	fmt.Scan(&valueA)

	for valueA > 0 {
		digito := valueA % 10
		invertido = invertido*10 + digito
		valueA /= 10
	}

	fmt.Println(invertido)
}
