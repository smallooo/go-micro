package main

import (
	"log-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type JSONMessage struct {
	Message string `json:"message"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// read json into var
	var requestPayload JSONPayload
	_ = app.readJSON(w, r, &requestPayload)

	// insert data
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, resp)

}

//-------------------------------------------message----------------------------------------

func (app *Config) Message(w http.ResponseWriter, r *http.Request) {
	// read json into var
	var jSONMessage JSONMessage
	_ = app.readJSON(w, r, &jSONMessage)

	// insert data
	event := data.MessageEntry{
		ID:   jSONMessage.Message,
		Data: jSONMessage.Message,
	}

	err := app.MessageModels.MessageEntry.InsertMessage(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}

	app.writeJSON(w, http.StatusAccepted, resp)

}

//------------------------------------------wechat--------------------------------------------
