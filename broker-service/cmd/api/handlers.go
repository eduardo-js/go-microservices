package main

import (
	"broker-service/event"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
	Mail   MailPayload `json:"mail,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
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
	case "log":
		app.logPayloadRPC(w, requestPayload.Log)
	case "mail":
		app.sendEmail(w, requestPayload.Mail)

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

	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	}
	if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error on auth service request"))
		return
	}
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

func (app *Config) logPayload(w http.ResponseWriter, entry LogPayload) {
	err := app.pushToQueue(entry.Name, entry.Data)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var resp jsonResponse
	resp.Error = false
	resp.Message = "Authenticated"
	resp.Data = "Logged successfully via rabbitMQ"
	app.writeJSON(w, http.StatusAccepted, resp)
}

type RPCPayload struct {
	Name string
	Data string
}

func (app *Config) logPayloadRPC(w http.ResponseWriter, entry LogPayload) {
	// err := app.pushToQueue(entry.Name, entry.Data)
	// if err != nil {
	// 	app.errorJSON(w, err)
	// 	return
	// }
	// var resp jsonResponse
	// resp.Error = false
	// resp.Message = "Authenticated"
	// resp.Data = "Logged successfully via rabbitMQ"
	// app.writeJSON(w, http.StatusAccepted, resp)
	client, err := rpc.Dial("tcp", "logger-service:5001")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var result string
	err = client.Call("RPCServer.LogInfo", RPCPayload(entry), &result)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var resp jsonResponse
	resp.Error = false
	resp.Message = result
	app.writeJSON(w, http.StatusAccepted, resp)
}

func (app *Config) sendEmail(w http.ResponseWriter, msg MailPayload) {
	jsonData, _ := json.MarshalIndent(msg, "", "\t")
	mailServiceURL := "http://mailer-service/send"
	request, err := http.NewRequest("POST", mailServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	request.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling mailer service"))
		return
	}
	var payload jsonResponse
	payload.Error = false
	payload.Message = "Message sent to" + msg.To
	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) pushToQueue(name, msg string) error {
	emitter, err := event.NewEventEmitter(app.RabbitMQ)
	if err != nil {
		return err
	}
	payload := LogPayload{
		Name: name,
		Data: msg,
	}
	j, _ := json.MarshalIndent(&payload, "", "\t")
	err = emitter.Push(string(j), "log.INFO")
	if err != nil {
		return err
	}
	return nil
}
