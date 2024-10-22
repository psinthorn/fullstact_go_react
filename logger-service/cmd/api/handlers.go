package main

import (
	"fmt"
	"net/http"
)

func (app *Config) Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Welcme to Logger")
}

func (app *Config) NewLog(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Create new log")
}
