package main

import "fmt"

type Transporte interface {
	Entregar()
}

type Moto struct {
	Placa string
}

type Caminhao struct {
	CargaMax int
}

func (moto Moto) Entregar() {
	fmt.Println("Moto placa", moto.Placa, "entregando pedido rapido")
}

func (caminhao Caminhao) Entregar() {
	fmt.Println("Caminhao com capacidade ", caminhao.CargaMax, " entregando carga pesada")
}

func GetTransporte(tipo string) Transporte {

	if tipo == "moto" {
		return Moto{Placa: "ABC-1234"}
	} else if tipo == "Caminhao" {
		return Caminhao{CargaMax: 10}
	}

	return nil
}

func main() {
	transport := GetTransporte("moto")
	transport.Entregar()

	transportPesado := GetTransporte("Caminhao")
	transportPesado.Entregar()
}
