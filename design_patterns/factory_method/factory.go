package main

import "fmt"

type Enemy interface {
	Attack()
}

type Soldier struct{}

type Mage struct{}

func (soldier Soldier) Attack() {
	fmt.Println("Soldado ataca com espada!")
}

func (mage Mage) Attack() {
	fmt.Println("Mago lança bola de fogo!")
}

func CreateEnemy(class string) Enemy {

	if class == "soldier" {
		return Soldier{}
	} else if class == "mage" {
		return Mage{}
	}

	return nil
}

func main() {
	attack := CreateEnemy("mage")
	attack.Attack()
}
