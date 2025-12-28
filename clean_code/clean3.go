package main

import "fmt"

type (
	Notification interface {
		Send()
	}

	EmailSender struct {
		Email string
	}
)

func (e *EmailSender) Send() {
	fmt.Println("Send Email:", e.Email)
}

func FinalizeOrder(orderId int, notifier Notification) {
	fmt.Println("Order completed", orderId)
	notifier.Send()
}

func main() {

	notifier := &EmailSender{
		Email: "Your order has been shipped!",
	}

	FinalizeOrder(10, notifier)
}
