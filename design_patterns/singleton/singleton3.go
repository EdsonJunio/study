package main

import (
	"fmt"
	"sync"
)

type Databases struct {
	UrlConexao string
}

var (
	instanc *Databases
	onc     sync.Once
)

func GetDatabases() *Databases {
	onc.Do(func() {
		instanc = &Databases{
			UrlConexao: "postgres://localhost:5432",
		}
	})
	return instanc
}

func main() {
	database := GetDatabases()
	fmt.Println(database)
}
