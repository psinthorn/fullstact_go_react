package main

import (
	"fmt"
	"log"
	"net/http"
)

// infoLog struct to hold request log data
type infoLog struct {
	RemoteAddr string
	Proto      string
	Method     string
	URI        string
	UserAgent  string
	Referer    string
}

// log method to print the log data
func (log *infoLog) log() {
	fmt.Printf("RemoteAddr: %s, Proto: %s, Method: %s, URI: %s, User-Agent: %s, Referer: %s\n",
		log.RemoteAddr,
		log.Proto,
		log.Method,
		log.URI,
		log.UserAgent,
		log.Referer,
	)
}

// Config struct to hold application configuration
type Config struct {
	InfoLog *log.Logger
}

// logRequest middleware to log request data
func (app *Config) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create an infoLog instance with request data
		logData := infoLog{
			RemoteAddr: r.RemoteAddr,
			Proto:      r.Proto,
			Method:     r.Method,
			URI:        r.URL.RequestURI(),
			UserAgent:  r.Header.Get("User-Agent"),
			Referer:    r.Header.Get("Referer"),
		}

		// Log the request data
		logData.log()

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
