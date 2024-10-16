package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct{}

const srvPort = "80"

func main() {

	app := Config{}

	log.Printf("Broker Service Starting on Port: %s\n ", srvPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", srvPort),
		Handler: app.routes(),
	}

	// start http server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
