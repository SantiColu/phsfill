package main

import (
	"math/rand"
	"time"

	"github.com/SantiColu/phsfill/api"
	"github.com/SantiColu/phsfill/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	a := api.Api{
		Url: "https://example.com/api/auth/login",
	}

	a.LoginUser(models.NewRandomUser())
}
