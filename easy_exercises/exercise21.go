package easy_exercises

import "fmt"

func main() {
	var word string
	var sum int
	vowels := "aeiouAEIOU"

	fmt.Scan(&word)

	for i := 0; i < len(word); i++ {
		for v := 0; v < len(vowels); v++ {
			if word[i] == vowels[v] {
				sum++
			}
		}
	}

	fmt.Println(sum)
}
