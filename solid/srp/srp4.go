package srp

import "fmt"

type (
	Employee struct {
		ID     int
		Name   string
		Salary float64
	}

	PaymentCalculator struct{}

	PaymentRepository struct{}

	EmailNotifier struct{}

	PaymentService struct {
		Calculator PaymentCalculator
		Repo       PaymentRepository
		Notifier   EmailNotifier
	}
)

func (PaymentCalculator) CalculateTax(salary float64) float64 {
	tax := salary * 0.20
	return salary - tax
}

func (PaymentRepository) SavingPayment(id int, netPay float64) {
	fmt.Printf("INSERT INTO payments (emp_id, amount) VALUES (%d, %.2f)\n", id, netPay)
}

func (EmailNotifier) SendEmail(name string, netPay float64) {
	fmt.Printf("Enviando email para %s: Você recebeu %.2f\n", name, netPay)
}

func (p PaymentService) Pay(e *Employee) {
	netPay := p.Calculator.CalculateTax(e.Salary)
	p.Repo.SavingPayment(e.ID, netPay)
	p.Notifier.SendEmail(e.Name, netPay)
}

func main() {
	employee := Employee{
		ID:     1,
		Name:   "Edson",
		Salary: 5000,
	}

	service := PaymentService{
		Calculator: PaymentCalculator{},
		Repo:       PaymentRepository{},
		Notifier:   EmailNotifier{},
	}

	service.Pay(&employee)
}
