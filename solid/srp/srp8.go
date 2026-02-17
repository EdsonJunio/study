package srp

import "fmt"

type (
	Userr struct {
		ID    int
		Name  string
		Email string
	}

	ValidatorName  struct{}
	ImageResizer   struct{}
	StorageService struct{}
	UserRepository struct{}
	EmailNotifierr struct{}

	ProfileOrchestrator struct {
		vali  ValidatorName
		img   ImageResizer
		store StorageService
		repo  UserRepository
		email EmailNotifier
	}
)

func (v *ValidatorName) Validate(u *Userr, novoNome string) error {
	if novoNome == "" {
		return fmt.Errorf("nome vazio")
	}

	u.Name = novoNome
	return nil
}

func (i *ImageResizer) Resize(novaFoto []byte) ([]byte, error) {
	if novaFoto == nil {
		return nil, fmt.Errorf("foto inválida")
	}

	fmt.Println("Redimensionando imagem para 500x500...")
	return novaFoto, nil
}

func (s *StorageService) Upload(foto []byte) error {
	fmt.Printf("Enviando %d bytes para o S3 (Bucket: users)...\n", len(foto))
	return nil
}

func (r *UserRepository) Update(u *Userr) error {
	fmt.Printf("UPDATE users SET name='%s' WHERE id=%d\n", u.Name, u.ID)
	return nil
}

func (e *EmailNotifier) Notify(u *Userr) error {
	fmt.Printf("Enviando email para %s: 'Seu perfil foi alterado'\n", u.Email)
	return nil
}

func (p *ProfileOrchestrator) UpdateProfile(u *Userr, novoNome string, novaFoto []byte) error {

	if err := p.vali.Validate(u, novoNome); err != nil {
		return err
	}

	img, err := p.img.Resize(novaFoto)
	if err != nil {
		return err
	}

	if err := p.store.Upload(img); err != nil {
		return err
	}

	if err := p.repo.Update(u); err != nil {
		return err
	}

	if err := p.email.Notify(u); err != nil {
		fmt.Println("Falha ao enviar email:", err)
	}

	return nil
}

func main() {

	user := &Userr{ID: 1, Name: "Edson", Email: "edson@gmail.com"}
	foto := []byte("dados-binarios-da-foto")

	orq := ProfileOrchestrator{
		vali:  ValidatorName{},
		img:   ImageResizer{},
		store: StorageService{},
		repo:  UserRepository{},
		email: EmailNotifier{},
	}

	if err := orq.UpdateProfile(user, "Edson Updated", foto); err != nil {
		fmt.Println("Erro:", err)
	}
}
