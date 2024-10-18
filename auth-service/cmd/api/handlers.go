package main

import "net/http"

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

}
