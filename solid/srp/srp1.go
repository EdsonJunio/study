package srp

import "fmt"

type TaxCalculator interface {
	Calculator(amount float64) float64
}

type DefaultTaxCalculator struct {
}

func (d *DefaultTaxCalculator) Calculator(amount float64) float64 {
	return amount * 0.15
}

type Notifier interface {
	Send(userID int, message string) error
}

type EmailNotifier struct {
}

func (e *EmailNotifier) Send(userID int, message string) error {
	fmt.Println("Enviando email para servidor SMTP...", message)
	return nil
}

type InvoiceService struct {
	taxCalculator TaxCalculator
	notifier      Notifier
}

func (i *InvoiceService) GenerateInvoice(userID int, amount float64) error {
	tax := i.taxCalculator.Calculator(amount)
	total := amount + tax

	message := fmt.Sprintf("Olá User %d, sua fatura é %.2f", userID, total)

	return i.notifier.Send(userID, message)

}
