package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/psinthorn/microservice_fullstack_golang_react_nextjs/auth-service/data"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const srvPort = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

	// Todo
	// 1. Connect to database
	// 2. server config and start

	// 1.1 Connect to database
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Database")
	}

	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	// 2.1 Server config and start
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", srvPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

// Implement to Connct to database
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {

		DbConnection, err := openDB(dsn)

		if err != nil {
			log.Println("Postgres database server is not ready ...")
			counts++

		} else {
			log.Println("Database is connected and your are ready to go :)")
			return DbConnection

		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Can't connect to Database will try to re-connect again")
		time.Sleep(time.Second * 2)
		continue
	}

}
