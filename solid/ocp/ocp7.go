package ocp

import (
	"errors"
	"fmt"
)

type (
	Order struct {
		Weight float64
	}

	Transporter interface {
		Calculate(o *Order) (float64, error)
	}

	PostOffice struct{}
	Jadlog     struct{}
	Loggi      struct{}

	ServiceTransporter struct{}
)

func (p *PostOffice) Calculate(o *Order) (float64, error) {
	if o.Weight <= 0 {
		return 0, errors.New(" Order cannot be zero or negative")
	}
	return 10.0 + (o.Weight * 2.0), nil
}

func (j *Jadlog) Calculate(o *Order) (float64, error) {
	if o.Weight <= 0 {
		return 0, errors.New(" Order cannot be zero or negative")
	}

	return 15.0 + (o.Weight * 1.0), nil
}

func (j *Loggi) Calculate(o *Order) (float64, error) {
	if o.Weight <= 0 {
		return 0, errors.New(" Order cannot be zero or negative")
	}

	return o.Weight * 5.0, nil
}

func (c *ServiceTransporter) Calculate(o *Order, t Transporter) (float64, error) {
	weight, err := t.Calculate(o)

	return weight, err

}

func main() {
	order := &Order{
		Weight: 15.00,
	}

	servicetransporter := ServiceTransporter{}

	weight, err := servicetransporter.Calculate(order, &Loggi{})
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(weight)
}
