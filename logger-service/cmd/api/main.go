package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/gommon/log"
)

type Config struct {
	DB *sql.DB
}

func main() {

	app := Config{}

	// fmt.Println("Logger service is starting from now on... ;)")

	// connect to database

	// set up http server
	const serverPort = "80"

	log.Printf("Logger service is staring on port: %s", serverPort)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
