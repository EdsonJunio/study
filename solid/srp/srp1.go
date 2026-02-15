package srp

import "fmt"

type (
	Funcionario struct {
		Nome    string
		Cargo   string
		Salario float64
	}

	CalculadoraDeBonus struct{}

	RepositorioFuncionario struct{}

	NotificadorEmail struct{}

	OrquestradorContratacao struct {
		calculadora CalculadoraDeBonus
		repositorio RepositorioFuncionario
		mensagem    NotificadorEmail
	}
)

func (c *CalculadoraDeBonus) Calcular(f *Funcionario) float64 {
	if f.Cargo == "Gerente" {
		return f.Salario * 0.2
	}
	return 0.0
}

func (r *RepositorioFuncionario) Salvar(f *Funcionario) error {
	fmt.Printf("SQL: INSERT INTO funcionarios VALUES ('%s', %.2f)\n", f.Nome, f.Salario)
	return nil
}

func (n *NotificadorEmail) EnviarBoasVindas(f *Funcionario) {
	fmt.Printf("Enviando email para %s...", f.Nome)
}

func (o *OrquestradorContratacao) Contratar(f *Funcionario) {

	bonus := o.calculadora.Calcular(f)
	f.Salario += bonus

	err := o.repositorio.Salvar(f)
	if err != nil {

		fmt.Println("Erro ao salvar")

	}

	o.mensagem.EnviarBoasVindas(f)
}

func main() {
	func1 := &Funcionario{
		Nome:    "Edson",
		Cargo:   "Gerente",
		Salario: 100.00,
	}

	calc := CalculadoraDeBonus{}
	repo := RepositorioFuncionario{}
	email := NotificadorEmail{}

	sistema := OrquestradorContratacao{
		calculadora: calc,
		repositorio: repo,
		mensagem:    email,
	}

	sistema.Contratar(func1)
}
