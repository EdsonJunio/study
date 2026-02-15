package srp

import (
	"fmt"
	"strings"
)

type (
	User struct {
		Name     string
		Email    string
		Password string
	}

	ValidationUser struct{}

	EmailValidator struct{}

	Persistence struct{}

	OrchestratorRegistration struct {
		Validator      ValidationUser
		EmailValidator EmailValidator
		persist        Persistence
	}
)

func (v *ValidationUser) Validation(user *User) error {
	if len(user.Password) < 8 {
		return fmt.Errorf("password too short")
	}

	return nil
}

func (e *EmailValidator) Email(user *User) error {
	if !strings.Contains(user.Email, "@") {
		return fmt.Errorf("email invalido")
	}

	return nil
}

func (p *Persistence) SaveUser(user *User) error {
	fmt.Printf("Saving user %s in the database...\n", user.Name)

	return nil
}

func (o *OrchestratorRegistration) OrchestratorServer(user *User) error {
	validator := o.Validator.Validation(user)
	if err := validator; err != nil {
		fmt.Println(err)
		return err
	}

	validatorEmail := o.EmailValidator.Email(user)
	if err := validatorEmail; err != nil {
		fmt.Println(err)
		return err
	}

	persist := o.persist.SaveUser(user)
	if err := persist; err != nil {
		fmt.Printf("Error saving to the database")
	}

	return nil
}

func main() {
	users := &User{
		Name:     "Edson",
		Email:    "edsonbh@gmail.com",
		Password: "testeeeee",
	}

	passWord := ValidationUser{}
	Email := EmailValidator{}
	repo := Persistence{}

	server := OrchestratorRegistration{
		Validator:      passWord,
		EmailValidator: Email,
		persist:        repo,
	}

	server.OrchestratorServer(users)
}
