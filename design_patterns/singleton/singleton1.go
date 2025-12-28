package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Qtd int
}

var (
	CounterInstance *Counter
	once            sync.Once
)

func GetCounter() *Counter {
	once.Do(func() {
		CounterInstance = &Counter{}
	})
	return CounterInstance
}

func (c *Counter) add() {
	c.Qtd++
}

func main() {

	c1 := GetCounter()
	c1.add()
	c1.add()
	c1.add()

	c2 := GetCounter()
	c2.add()
	c2.add()
	c2.add()

	fmt.Println(c1.Qtd)
}
