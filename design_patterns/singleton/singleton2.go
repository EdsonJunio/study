package main

import (
	"fmt"
	"sync"
)

type Config struct {
	AppMode string
}

var (
	config *Config
	onces  sync.Once
)

func GetConfig() *Config {
	onces.Do(func() {
		config = &Config{}
	})
	return config
}

func main() {
	c1 := GetConfig()
	c1.AppMode = "Dark"
	c2 := GetConfig()
	fmt.Println(c2.AppMode)

}
