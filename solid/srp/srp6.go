package srp

import (
	"fmt"
	"io"
	"net/http"
)

type (
	ProductPhysical struct {
		Name   string
		Weight float64
	}

	CalculatorBase struct{}

	ServiceCEP struct{}

	OrchestratorPrice struct {
		calc CalculatorBase
		cep  ServiceCEP
	}
)

const PRICE_BASE = 0.0

func (s *ServiceCEP) GetUF(cep string) (string, error) {
	resp, err := http.Get("http://api-ficticia-correios.com/" + cep)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Resposta API:", string(body))

	return "SP", nil
}

func (c *CalculatorBase) Calculator(uf string, p *ProductPhysical) float64 {
	priceBase := PRICE_BASE

	if uf == "SP" {
		priceBase = 10.0
	} else {
		priceBase = 20.0
	}

	totalCost := priceBase * p.Weight
	return totalCost
}

func (o *OrchestratorPrice) CalculateTotal(cep string, p *ProductPhysical) (float64, error) {

	uf, err := o.cep.GetUF(cep)
	if err != nil {
		return 0, err
	}

	baseCost := o.calc.Calculator(uf, p)

	return baseCost, nil
}

func main() {
	prod := &ProductPhysical{
		Name:   "Geladeira",
		Weight: 60.0,
	}

	orq := OrchestratorPrice{
		calc: CalculatorBase{},
		cep:  ServiceCEP{},
	}

	valor, err := orq.CalculateTotal("01001-000", prod)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Printf("Frete total: R$ %.2f\n", valor)
}
