package main

import (
	"fmt"
)

func main() {
	text := "amor"

	for i := len(text) - 1; i >= 0; i-- {
		fmt.Println(string(text[i]))
	}
}
