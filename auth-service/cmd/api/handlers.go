package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {

	// Get user name and password data from request
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Read json data from request payload and check it error or not
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
	}

	//Compare user and password
	// 1. validate user againts database by email address
	fmt.Println(requestPayload.Email)
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid email address"), http.StatusUnauthorized)
		return

	}

	// Compare password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid password"), http.StatusUnauthorized)
		return
	}

	// if not error then prepare data for write out back to request
	// not need to check error here we have been check from above already
	payload := jsonResponse{
		Message: fmt.Sprintf("User %s is logged in", user.Email),
		Data:    user,
		Error:   false,
	}

	app.writeJSON(w, http.StatusAccepted, payload)

}
