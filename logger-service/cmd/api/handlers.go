package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/psinthorn/microservice_fullstack_golang_react_nextjs/data"
)

type JSONPayload struct {
	Name      string `json:"name"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Data      string `json:"data"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`
}

func (app *Config) Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Welcme to Logger")
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Write new log")
	// what we need to keep
	// log level `{"level":"info","message":"hello world"}`
	// collect the log level and message from the request to variable
	requestPayload := JSONPayload{}
	_ = app.readJSON(w, r, &requestPayload)

	// validate the log level
	// if the log level is not info, warn, or error, return an error
	// if requestPayload.Level != "info" && requestPayload.Level != "warn" && requestPayload.Level != "error" {
	// 	app.writeJSON(w, http.StatusBadRequest, jsonResponse{
	// 		Error:   true,
	// 		Message: "Invalid log level",
	// 	})
	// 	return
	// }

	// write to a database
	// insert the log level and message into the database
	// if there is an error, return an error
	logEntry := data.LogEntry{
		Name:      requestPayload.Name,
		Data:      requestPayload.Data,
		Level:     requestPayload.Level,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := app.Models.LogEntry.Insert(logEntry)
	if err != nil {
		app.writeJSON(w, http.StatusInternalServerError, jsonResponse{
			Error:   true,
			Message: "Failed to write log",
		})
		return
	}

	// return the log level and message

	response := jsonResponse{
		Error:   false,
		Message: "Log written",
		Data:    logEntry,
	}

	app.writeJSON(w, http.StatusOK, response)

}
