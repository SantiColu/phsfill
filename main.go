package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/SantiColu/phsfill/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// a := api.Api{
	// 	Url: "https://example.com/api/auth/login",
	// }

	for i := 0; i < 15; i++ {
		u := models.NewRandomUser()
		fmt.Printf("Nombre: %v, Apellido: %v, Email: %v, Password: %v\n", u.Name, u.Lastname, u.Email, u.Password)
	}

	// a.LoginUser(models.NewRandomUser())
}
