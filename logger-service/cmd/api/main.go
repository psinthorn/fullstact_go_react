package main

import (
	"database/sql"
	"fmt"
)

type Config struct {
	DB *sql.DB
}

func main() {

	app := Config{}

	fmt.Println("Logger service is starting from now on... ;)")

	// connect to database

	// set up http server
	const serverPort = "80"

}
