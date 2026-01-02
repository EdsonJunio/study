package main

import (
	"fmt"
	"sync"
	"time"
)

type Config struct {
	Version string
	mu      sync.Mutex
}

var (
	configInstance *Config
	onces          sync.Once
)

func GetConfig() *Config {
	onces.Do(func() {
		configInstance = &Config{
			Version: "1.0",
		}
	})
	return configInstance
}

func (c *Config) SetVersion(v string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	fmt.Println("Atualizando versão para:", v)
	time.Sleep(time.Millisecond * 500)
	c.Version = v
}

func main() {
	config := GetConfig()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		config.SetVersion("2.0")
	}()

	go func() {
		defer wg.Done()
		config.SetVersion("3.0")
	}()

	wg.Wait()

	fmt.Println("Versão final:", config.Version)
}
