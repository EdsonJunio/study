package easy_exercises

import "fmt"

func main() {
	var valueA int

	fmt.Scan(&valueA)

	if valueA >= 10 && valueA <= 20 || valueA >= 100 && valueA <= 200 {
		fmt.Println("Seguro")
		return
	}

	fmt.Println("Inseguro")
}
