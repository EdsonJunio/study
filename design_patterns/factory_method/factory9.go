package main

import (
	"errors"
	"fmt"
)

type (
	Button interface {
		Render()
	}

	ButtonWindows struct{}

	ButtonLinux struct{}
)

func (*ButtonWindows) Render() {
	fmt.Printf("Imprime [WIN] ----------------- [WIN]")
}

func (*ButtonLinux) Render() {
	fmt.Printf("LIN) ################# (LIN")
}

func CreateButton(operationalSystem string, active bool) (Button, error) {
	if operationalSystem == "" || !active {
		return nil, errors.New("ACTIVE_LICENSE false")
	}

	switch operationalSystem {
	case "windows":
		return &ButtonWindows{}, nil
	case "linux":
		return &ButtonLinux{}, nil
	default:
		return nil, errors.New("unsupported operational system")
	}
}

func main() {
	const ACTIVE_LICENSE_WINDOWS = false
	const ACTIVE_LICENSE_LINUX = true

	systemW, err := CreateButton("windows", ACTIVE_LICENSE_WINDOWS)
	if err != nil {
		fmt.Println(err)
	} else {
		systemW.Render()
	}

	systemL, err := CreateButton("linux", ACTIVE_LICENSE_LINUX)
	if err != nil {
		fmt.Println(err)
	} else {
		systemL.Render()
	}
}
