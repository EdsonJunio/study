package main

import "fmt"

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBNam    string
}

func ConnectToDatabase(dsn *Config) {
	fmt.Println("Conectando em:", dsn)

}

func main() {
	config := &Config{
		Host:     "localhost",
		Port:     8080,
		User:     "edson",
		Password: "key",
		DBNam:    "ecommerce",
	}

	ConnectToDatabase(config)
}
