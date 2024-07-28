package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	//
	var requestPayload struct {
		email    string `json:"email"`
		password string `json:"password"`
	}
	err := app.ReadJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the request payload against the database
	user, err := app.Models.User.GetByEmail(requestPayload.email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid Credentials"), http.StatusUnauthorized)
		return
	}
	valid, err := user.PasswordMatches(requestPayload.password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid Credentials"), http.StatusUnauthorized)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("logged in as %s", user.Email),
		Data:    user,
	}

	app.WriteJSON(w, http.StatusAccepted, payload)
}
