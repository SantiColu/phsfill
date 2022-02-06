package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/SantiColu/phsfill/models"
)

type register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Api struct {
	Url      string
	WaitTime time.Duration
}

func (api *Api) LoginUser(u *models.User) {
	data := login{
		Email:    u.Email,
		Password: u.Password,
	}
	api.SendJSON(data)
}

func (api *Api) LoginUserForm(u *models.User) {
	data := url.Values{"email": {u.Email}, "password": {u.Password}}
	api.SendForm(data)
}

func (api *Api) LoginUserBasic(u *models.User) {
	data := login{
		Email:    u.Email,
		Password: u.Password,
	}

	reqBody, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("[i] Sending %v\n", string(reqBody))

	client := &http.Client{}
	req, err := http.NewRequest("POST", api.Url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.SetBasicAuth(u.Email, u.Password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(resp)
		return
	}

	fmt.Printf("[i] Server responded with %v\n\n", resp.StatusCode)
}

func (api *Api) RegisterUser(u *models.User) {
	data := register{
		Email:    u.Email,
		Password: u.Password,
		Name:     u.Name,
		Lastname: u.Lastname,
	}

	api.SendJSON(data)
}

func (api *Api) RegisterUserForm(u *models.User) {
	data := url.Values{"email": {u.Email}, "password": {u.Password}, "name": {u.Name}, "lastname": {u.Lastname}}
	api.SendForm(data)
}

func (api *Api) SendForm(data url.Values) {

	reqBody, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("[i] Sending %v\n", string(reqBody))

	resp, err := http.PostForm(api.Url, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[i] Server responded with %v\n\n", resp.StatusCode)
}

func (api *Api) SendJSON(data interface{}) {
	reqBody, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("[i] Sending %v\n", string(reqBody))

	resp, err := http.Post(api.Url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("[i] Server responded with %v\n\n", resp.StatusCode)
}
