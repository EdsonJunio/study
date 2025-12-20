package main

import (
	"errors"
	"fmt"
)

type (
	Parser interface {
		Parse(content []byte)
	}

	JsonParser struct{}

	XmlParser struct{}
)

func (*JsonParser) Parse(content []byte) {
	fmt.Printf("Decoding JSON: %s \n", content)
}

func (*XmlParser) Parse(content []byte) {
	fmt.Printf("Decoding tags XML: %d \n", content)
}

func GetParser(fileName string) (Parser, error) {
	switch fileName {
	case "dados.json":
		return &JsonParser{}, nil
	case "config.xml":
		return &XmlParser{}, nil
	default:
		return nil, errors.New("invalid processor content")
	}

}

func main() {
	files := []string{"dados.json", "config.xml", "foto.png"}

	for _, file := range files {
		parser, err := GetParser(file)
		if err != nil {
			fmt.Println(err)
			continue
		}

		parser.Parse([]byte("conteudo fake"))
	}
}
