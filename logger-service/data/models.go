package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name" json:"name"`
	Data      string    `bson:"data" json:"data"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func New(mongo *mongo.Client) Models {
	return Models{
		LogEntry: LogEntry{},
	}
}

func (l *LogEntry) Insert(logEntry LogEntry) error {
	// mongo context connection and declair collection
	collection := client.Database("logs").Collection("logs")

	// insert Log to collection
	_, err := collection.InsertOne(context.TODO(), LogEntry{
		Name:      logEntry.Name,
		Data:      logEntry.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("error log insert: ", err)
		return err
	}

	//  nil if no error
	return nil
}

// Get and show all cuurent logs
func (l *LogEntry) AllLogs() ([]*LogEntry, error) {
	// declair variable
	var logs []*LogEntry

	// Create context for time out and defer
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// select database and collection
	collection := client.Database("logs").Collection("logs")

	// declaire options method when find all logs
	opts := options.Find()
	// sort logs by created_at by DESC
	opts.SetSort(bson.D{{"created_at", -1}})

	// get all logs
	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("Error to find all logs: ", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var logItem LogEntry
		err := cursor.Decode(&logItem)
		if err != nil {
			log.Println("decode log to slice of logs is error: ", err)
			return nil, err
		} else {
			logs = append(logs, &logItem)
		}
	}

	return logs, nil

}
