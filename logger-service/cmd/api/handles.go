package main

import (
	"logger-service/data"
	"net/http"

	"github.com/tsawler/toolbox"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

var t toolbox.Tools

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// read json into var
	var requestPayload JSONPayload
	_ = t.ReadJSON(w, r, &requestPayload)
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}
	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}
	resp := toolbox.JSONResponse{
		Error:   false,
		Message: "logged",
	}
	t.WriteJSON(w, http.StatusAccepted, resp)
}
