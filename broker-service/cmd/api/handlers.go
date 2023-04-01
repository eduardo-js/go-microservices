package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hello from Broker",
	}
	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleRequest(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	switch requestPayload.Action {
	case "auth":
		app.authenticate(w, requestPayload.Auth)

	default:
		app.errorJSON(w, errors.New("invalid action"))
	}
}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	data, _ := json.MarshalIndent(a, "", "\t")
	request, err := http.NewRequest("POST", "http://auth-service/authenticate", bytes.NewBuffer(data))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	fmt.Println("b")
	client := &http.Client{}
	response, err := client.Do(request)
	fmt.Println(response, err)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	fmt.Println("d")
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	}
	fmt.Println("e")
	if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error on auth service request"))
		return
	}
	fmt.Println("f")

	var jsonFromService jsonResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
	}

	var resp jsonResponse
	resp.Error = false
	resp.Message = "Authenticated"
	resp.Data = jsonFromService.Data
	app.writeJSON(w, http.StatusAccepted, resp)
}
