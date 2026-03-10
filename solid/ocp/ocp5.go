package ocp

import (
	"errors"
	"fmt"
)

type (
	Employee struct {
		Name   string
		Salary float64
	}

	Nation interface {
		Calculate(e *Employee) (float64, error)
	}

	Brazil   struct{}
	EUA      struct{}
	Portugal struct{}

	CalculatorService struct{}
)

func (b *Brazil) Calculate(e *Employee) (float64, error) {

	if e.Salary <= 0 {
		return 0, errors.New("salary cannot be zero or negative")
	}
	salary := e.Salary * 0.27

	return salary, nil
}

func (eua *EUA) Calculate(e *Employee) (float64, error) {

	if e.Salary <= 0 {
		return 0, errors.New("salary cannot be zero or negative")
	}
	salary := e.Salary * 0.35

	return salary, nil
}

func (p *Portugal) Calculate(e *Employee) (float64, error) {

	if e.Salary <= 0 {
		return 0, errors.New("salary cannot be zero or negative")
	}
	salary := e.Salary * 0.45

	return salary, nil
}

func (p *CalculatorService) Calculate(e *Employee, n Nation) (float64, error) {
	salary, err := n.Calculate(e)
	return salary, err
}

func main() {
	employee := &Employee{
		Name:   "Edson",
		Salary: 4500.00,
	}

	service := CalculatorService{}

	salary, err := service.Calculate(employee, &EUA{})
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(salary)

}
