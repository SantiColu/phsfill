package models

import (
	"math/rand"

	"github.com/SantiColu/phsfill/utils"
)

type User struct {
	Name     string
	Lastname string
	Password string
	Email    string
}

func NewUser(name, lastname, password, email string) *User {
	user := &User{name, lastname, password, email}
	return user
}

func NewRandomUser() *User {
	name := utils.GenerateName()
	lastname := utils.GenerateLastname()
	email := utils.GenerateEmail(name, lastname)
	password := utils.SelectCommonPassword()

	if rand.Intn(2) == 1 { // random email
		email = utils.GenerateRandomEmail()
	}

	if rand.Intn(3) == 1 { // random password
		password = utils.GenerateRandomPassword(rand.Intn(16) + 8)
	}

	user := &User{name, lastname, password, email}
	return user
}
