package ocp

import "fmt"

type (
	User struct {
		Email    string
		PassWord string
	}

	LoginGlobal interface {
		Login(u *User) bool
	}

	Local        struct{}
	Google       struct{}
	LoginService struct{}
)

func (l *Local) Login(u *User) bool {
	fmt.Printf("Searching for %s in the database...\n", u.Email)
	return true

}

func (g *Google) Login(u *User) bool {
	fmt.Printf("Validating Google token for %s...\n", u.Email)
	return true
}

func (l *LoginService) Login(u *User, lo LoginGlobal) bool {
	login := lo.Login(u)
	return login
}

func main() {
	login := &User{
		Email:    "edson@gmail.com",
		PassWord: "123",
	}

	serviceLogin := LoginService{}

	success := serviceLogin.Login(login, &Local{})
	if success != true {
		return
	}
}
