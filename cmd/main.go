package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"time"

	"github.com/SantiColu/phsfill/api"
	"github.com/SantiColu/phsfill/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	typePtr := flag.String("type", "login", "what action should the script do [valid options: login, register]")
	requestsPtr := flag.Int("requests", 500, "number of requests")
	delayPtr := flag.Int("delay", 1000, "delay between each request in milliseconds")
	formPtr := flag.Bool("form", false, "if it is true the data will be sended as a form")
	basicPtr := flag.Bool("basic", false, "if it is true the data will be sended with basic-auth format")
	urlPtr := flag.String("url", "", "the scammer site url, where the fake data will be sended")

	flag.Usage = usage

	flag.Parse()

	// CHECKS
	if *typePtr != "login" && *typePtr != "register" {
		fmt.Printf("type must be 'login' or 'register'\n\n")
		flag.Usage()
		os.Exit(2)
	}

	if _, err := url.ParseRequestURI(*urlPtr); err != nil {
		fmt.Printf("url must be valid, example: 'http://google.com/'\n\n")
		flag.Usage()
		os.Exit(2)
	}

	// INITIALIZATION
	a := api.Api{
		Url: *urlPtr,
	}

	sendFunc := func(u *models.User) {}

	if *typePtr == "login" {
		sendFunc = a.LoginUser

		if *formPtr {
			sendFunc = a.LoginUserForm
		}

		if *basicPtr {
			sendFunc = a.LoginUserBasic
		}
	}

	if *typePtr == "register" {
		sendFunc = a.RegisterUser

		if *formPtr {
			sendFunc = a.RegisterUserForm
		}
	}

	fmt.Println("\nConfiguration:")
	fmt.Println("Type:", *typePtr)
	fmt.Println("Requests:", *requestsPtr)
	fmt.Println("Delay:", *delayPtr)
	fmt.Println("Form:", *formPtr)
	fmt.Println("Basic:", *basicPtr)
	fmt.Println("URL:", *urlPtr)
	fmt.Printf("\n\n")

	for i := 0; i < *requestsPtr; i++ {
		sendFunc(models.NewRandomUser())
		time.Sleep(time.Duration(*delayPtr) * time.Millisecond)
	}
}

func usage() {
	fmt.Println("usage: phsfill [OPTIONS] \n\nphsfill is a simple tool fill phishing sites databases with random data and prevent theft")
	flag.PrintDefaults()
}
