package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"` // auth is store user and password from requst
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Message: "Broker is Alive",
		Error:   false,
	}

	out, _ := json.MarshalIndent(payload, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	// create variables
	var requestPayload RequestPayload

	// get data from the request into json format if error then return
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	// check the request action
	switch requestPayload.Action {
	case "auth":
		app.authenticate(w, requestPayload.Auth)

	default:
		app.errorJSON(w, errors.New("unknow action"))
	}

}

// Check authentication with auth payload
func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	// Prepare json data
	jsonDataFromRequest, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// send request to authentication url with json data for authenticate
	// use http client to send request
	// create new http request

	request, err := http.NewRequest("POST", "http://authentication-service/handle", bytes.NewBuffer(jsonDataFromRequest))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
	}
	defer response.Body.Close()

	// check the respone status code to make sure we can contact to authentication service
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid user name or password"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("can't contac to authentication service"))
		return
	}

	// if not error then decode data from authentication service json to variable
	var jsonDataFromAuthService jsonResponse
	err = json.NewDecoder(response.Body).Decode(&jsonDataFromAuthService)
	if err != nil {
		app.errorJSON(w, err, http.StatusUnauthorized)
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = jsonDataFromAuthService.Message
	payload.Data = jsonDataFromAuthService.Data

	// return data to request
	app.writeJSON(w, http.StatusAccepted, payload)

}
