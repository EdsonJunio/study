package srp

import (
	"fmt"
	"time"
)

type (
	Payment struct {
		Value float64
	}

	Logger struct{}

	ValidatorPayment struct{}

	PaymentOrchestrator struct {
		Log     Logger
		Payment ValidatorPayment
	}
)

func (l *Logger) LogStart(p *Payment) error {
	fmt.Printf("[LOG - %s]: Starting processing of R$%.2f\n", time.Now().Format(time.RFC3339), p.Value)
	return nil
}

func (v *ValidatorPayment) PaymentValidator(p *Payment) error {
	if p.Value > 0 {
		fmt.Println("Payment Successfully Completed")
	}

	return nil
}

func (l *Logger) LogFinal(p *Payment) error {
	fmt.Printf("[LOG - %s]: End of processing\n", time.Now().Format(time.RFC3339))
	return nil
}

func (pay *PaymentOrchestrator) ProcessPayment(Payment *Payment) error {
	loggerStart := pay.Log.LogStart(Payment)
	if err := loggerStart; err != nil {
		fmt.Println(err)
	}

	validatorPayment := pay.Payment.PaymentValidator(Payment)
	if err := validatorPayment; err != nil {
		fmt.Println(err)
		return err
	}

	loggerFinal := pay.Log.LogFinal(Payment)
	if err := loggerFinal; err != nil {
		fmt.Println(err)
	}

	return nil
}

func main() {
	payment := &Payment{
		Value: 100.00,
	}

	logs := Logger{}
	validator := ValidatorPayment{}

	process := PaymentOrchestrator{
		Log:     logs,
		Payment: validator,
	}

	process.ProcessPayment(payment)

}
