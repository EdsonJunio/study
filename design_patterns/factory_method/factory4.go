package main

import (
	"errors"
	"fmt"
)

type (
	Database interface {
		Conectar()
	}

	Postgres struct {
	}

	MySQL struct {
	}
)

func (Postgres) Conectar() {
	fmt.Println("Connect PostGres")
}

func (MySQL) Conectar() {
	fmt.Println("Connect MySQL")
}

func GetDatabase(tip string) (Database, error) {

	switch tip {
	case "postgres":
		return Postgres{}, nil
	case "mysql":
		return MySQL{}, nil
	default:
		return nil, errors.New("bank not supported\n")
	}

}

func main() {
	bank, err := GetDatabase("postgres")

	if err != nil {
		fmt.Println("error", err)
		return
	}

	bank.Conectar()
}
