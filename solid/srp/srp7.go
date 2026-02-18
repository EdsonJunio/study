package srp

import "fmt"

type (
	Conta struct {
		Saldo float64
	}

	RegraDeTransferencia struct{}

	Loggers struct{}

	AntFraude struct{}

	OrquestradorTranferencia struct {
		transf  RegraDeTransferencia
		loggers Loggers
		ant     AntFraude
	}
)

func (r *RegraDeTransferencia) Transferir(origem, destino *Conta, valor float64) error {

	if valor <= 0 {
		return fmt.Errorf("valor inválido")
	}

	if origem.Saldo < valor {
		return fmt.Errorf("saldo insuficiente")
	}

	origem.Saldo -= valor
	destino.Saldo += valor

	return nil
}

func (l *Loggers) Log() error {
	fmt.Println("[AUDITORIA] Transferência realizada com sucesso. Valor:")

	return fmt.Errorf("falha ao salvar log no serviço externo")
}

func (a *AntFraude) AlertAntfraude(valor float64) error {

	if valor > 10000 {
		return fmt.Errorf("[ANTI-FRAUDE] Transferência bloqueada por valor suspeito")
	}

	return nil
}

func (o *OrquestradorTranferencia) OrquestradorT(origem, destino *Conta, valor float64) error {

	if err := o.ant.AlertAntfraude(valor); err != nil {
		return err
	}

	if err := o.transf.Transferir(origem, destino, valor); err != nil {
		return err
	}

	if err := o.loggers.Log(); err != nil {
		fmt.Println("Aviso: transferência concluída, mas falhou auditoria:", err)
	}

	return nil
}

func main() {

	c1 := &Conta{Saldo: 60000}
	c2 := &Conta{Saldo: 20000}

	orq := OrquestradorTranferencia{
		transf:  RegraDeTransferencia{},
		loggers: Loggers{},
		ant:     AntFraude{},
	}

	err := orq.OrquestradorT(c1, c2, 500)

	if err != nil {
		fmt.Println("Falha na transferência:", err)
		return
	}

	fmt.Println("Transferência concluída!")
	fmt.Println("Saldo origem:", c1.Saldo)
	fmt.Println("Saldo destino:", c2.Saldo)
}
