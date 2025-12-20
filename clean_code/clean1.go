/*func GetUser(id int) User {
	if id <= 0 {
		return User{} // Erro de validação
	}
	// SQL misturado na regra
	query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", id)
	fmt.Println("Executando no banco:", query) // Simula o banco
	return User{ID: id, Name: "UserDB"}
}*/

package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type UserService struct {
	repo UserRepository
}

type UserRepository interface {
	GetByID(id int) (User, error)
}

type UserRepositoryDB struct{}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (User, error) {
	if id <= 0 {
		return User{}, fmt.Errorf("id inválido")
	}

	return s.repo.GetByID(id)
}

func (r *UserRepositoryDB) GetByID(id int) (User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", id)
	fmt.Println("Executando no banco:", query)

	return User{ID: id, Name: "UserDB"}, nil
}

func main() {
	repo := &UserRepositoryDB{}
	service := NewUserService(repo)

	user, err := service.GetUser(1)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Usuário:", user)
}
