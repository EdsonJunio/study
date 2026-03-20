package ocp

import "fmt"

type (
	Document struct {
		Number string
	}

	ValidatorDocument interface {
		Validate(d *Document) bool
	}

	ValidatorCPF      struct{}
	ValidatorCNPF     struct{}
	ValidatorPassport struct{}
	ValidatorService  struct{}
)

func (v *ValidatorCPF) Validate(d *Document) bool {
	fmt.Println("Running CPF validation algorithm for", d.Number)
	return len(d.Number) == 11
}

func (v *ValidatorCNPF) Validate(d *Document) bool {
	fmt.Println("Running CNPJ validation algorithm for", d.Number)
	return len(d.Number) == 14
}

func (v *ValidatorPassport) Validate(d *Document) bool {
	fmt.Println("Validating Passport Immigration Rules", d.Number)
	return len(d.Number) == 8
}

func (v *ValidatorService) Validate(d *Document, vs ValidatorDocument) bool {
	validator := vs.Validate(d)
	return validator
}

func main() {
	doc := &Document{
		Number: "12345678901234"}

	validator := ValidatorService{}

	valido := validator.Validate(doc, &ValidatorCNPF{})
	fmt.Println("Valid document?", valido)

	doc2 := &Document{
		Number: "12345678901"}

	validator2 := ValidatorService{}

	valido2 := validator2.Validate(doc2, &ValidatorCPF{})
	fmt.Println("Valid document?", valido2)
}
