package easy_exercises

import "fmt"

func main() {
	var text string
	fmt.Scan(&text)

	if len(text) >= 2 && text[0] == 'G' && text[1] == 'o' {
		fmt.Println("Válido")
		return
	}

	fmt.Println("Inválido")
}
