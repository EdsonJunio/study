package ocp

import "fmt"

type Pedido struct {
	ValorTotal float64
}

type Desconto interface {
	Calcular(p *Pedido) float64
}

type DescontoComum struct{}

func (d DescontoComum) Calcular(p *Pedido) float64 {
	return p.ValorTotal
}

type DescontoEstudante struct{}

func (d DescontoEstudante) Calcular(p *Pedido) float64 {
	return p.ValorTotal * 0.5
}

type DescontoVIP struct{}

func (d DescontoVIP) Calcular(p *Pedido) float64 {
	return p.ValorTotal * 0.8
}

type CalculadoraDeDesconto struct{}

func (c *CalculadoraDeDesconto) Calcular(p *Pedido, d Desconto) float64 {
	return d.Calcular(p)
}

func main() {
	pedido := &Pedido{
		ValorTotal: 100.0,
	}

	pedido2 := &Pedido{
		ValorTotal: 20.00,
	}
	calc := CalculadoraDeDesconto{}

	valorFinal := calc.Calcular(pedido, DescontoVIP{})
	fmt.Printf("Valor a pagar: R$ %.2f\n", valorFinal)

	valorFinal2 := calc.Calcular(pedido2, DescontoComum{})
	fmt.Printf("Valor a pagar: R$ %.2f\n", valorFinal2)
}
