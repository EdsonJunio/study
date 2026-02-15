package srp

import "fmt"

type (
	EmailSender struct {
		Host string
		Port int
	}

	OrchestratorEmail struct {
		Email EmailSender
	}
)

func (e *EmailSender) Send(msg string) error {
	fmt.Printf("Conectando em %s:%d para enviar: %s\n", e.Host, e.Port, msg)
	return nil
}

func (o *OrchestratorEmail) OrchestratorSend(msg string) error {
	return o.Email.Send(msg)
}

func main() {
	orchestratorGmail := OrchestratorEmail{
		Email: EmailSender{
			Host: "smtp.gmail.com",
			Port: 587,
		},
	}

	orchestratorLocal := OrchestratorEmail{
		Email: EmailSender{
			Host: "localhost",
			Port: 2525,
		},
	}

	orchestratorGmail.OrchestratorSend("Mensagem via Gmail")
	orchestratorLocal.OrchestratorSend("Mensagem via servidor Local")
}
