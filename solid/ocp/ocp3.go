package ocp

import "fmt"

type (
	ServiceNotification struct {
		Message string
	}

	Notification interface {
		Send(s *ServiceNotification) error
	}

	EmailNotification struct{}

	SMSNotification struct{}

	NotificationOrchestrator struct{}
)

func (e *EmailNotification) Send(s *ServiceNotification) error {
	fmt.Println("Sending E-mail corporation:", s.Message)
	return nil
}

func (sms *SMSNotification) Send(s *ServiceNotification) error {
	fmt.Println("Sending SMS via operator:", s.Message)
	return nil
}

func (no *NotificationOrchestrator) Send(s *ServiceNotification, n Notification) error {
	err := n.Send(s)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	notification := &ServiceNotification{
		Message: "Your transfer of R$ 5,000 has been approved",
	}

	not := NotificationOrchestrator{}
	err := not.Send(notification, &EmailNotification{})
	if err != nil {
		fmt.Println("Erro:", err)
	}

}
