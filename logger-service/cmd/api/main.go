package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// set up server serivce port
const (
	webPort  = "80"
	mongoURL = "mongodb://mongo:27017"
	rpcPort  = "5001"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
	DB *sql.DB
}

func main() {

	app := Config{}

	// connect to database
	client, err := connectToMongoDB()
	if err != nil {
		log.Panic(err)
	}

	// create context in case of disconnect
	ctx, cancle := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancle()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// http server start from here
	log.Printf("Logger service is staring on port: %s", webPort)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func connectToMongoDB() (*mongo.Client, error) {

	// create MongoDB clients options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "admin",
	})

	// create connection
	conn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Print("can't connect to MongoDB error: %s", err)
		return nil, err
	}

	return conn, nil
}
