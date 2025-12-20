package main

import "fmt"

type Logger interface {
	Log(mensage string)
}

type ConsoleLogger struct{}

type FileLogger struct{}

func (console ConsoleLogger) Log(mensage string) {
	fmt.Println("[CONSOLE]:", mensage)
}

func (file FileLogger) Log(mensage string) {
	fmt.Println("[ARQUIVO]: Salvando no disco...", mensage)
}

func GetLogger(tip string) Logger {

	if tip == "console" {
		return ConsoleLogger{}
	} else if tip == "file" {
		return FileLogger{}
	}

	return nil
}

func main() {

	configUsuario := "file"

	meuLogger := GetLogger(configUsuario)

	meuLogger.Log("O sistema iniciou...")
	meuLogger.Log("Conectando ao banco...")
}
