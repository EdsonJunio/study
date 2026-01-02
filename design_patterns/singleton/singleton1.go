package main

import (
	"fmt"
	"sync"
)

type Database struct {
}

var (
	instance *Database
	once     sync.Once
)

func GetDatabase() *Database {
	once.Do(func() {
		instance = &Database{}
	})
	return instance
}

func main() {
	db1 := GetDatabase()
	b2 := GetDatabase()

	fmt.Printf("endereço de memória %p \n", db1)
	fmt.Printf("endereço de memória %p \n ", b2)

	if db1 == b2 {
		fmt.Println("É igual!")
	} else {
		fmt.Println("É diferente!")
	}
}
