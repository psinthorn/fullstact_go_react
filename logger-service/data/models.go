package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        string    `bson:"_id,omitempty"json:"id,omitempty"`
	Name      string    `bson:"name"json:"name"`
	Data      string    `bson:"data"json:"data"`
	Level     string    `bson:"level,omitempty"json:"level,omitempty"`
	CreatedAt time.Time `bson:"created_at"json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"json:"updated_at"`
}

func New(mongo *mongo.Client) Models {
	return Models{
		LogEntry: LogEntry{},
	}
}

func (l *LogEntry) Insert(logEntry LogEntry) error {
	// mongo context connection and declair collection
	collection := client.Database("f2logs").Collection("logs")

	// insert Log to collection
	_, err := collection.InsertOne(context.TODO(), LogEntry{
		Name:      logEntry.Name,
		Data:      logEntry.Data,
		Level:     logEntry.Level,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("insert log error: ", err)
		return err
	}

	//  nil if no error
	return nil
}

// Get and show all cuurent logs
func (l *LogEntry) GetAll() ([]*LogEntry, error) {
	// declair variable
	var logs []*LogEntry

	// Create context for time out and defer
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// select database and collection
	collection := client.Database("f2logs").Collection("logs")

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

// GetById: get log by id
func (l *LogEntry) GetById(id string) (*LogEntry, error) {
	// declair varable
	var logEntry LogEntry

	// if can't connect to database within 15 seccond then cancle
	ctx, cancle := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancle()

	// select database and collection
	collection := client.Database("f2logs").Collection("logs")

	logId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": logId}).Decode(&logEntry)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &logEntry, nil
}

// DropCollection:
func (l *LogEntry) DropCollection() error {
	// Create context if can't connect to database within 15 seconds the cancle process
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Select collection to work with
	collection := client.Database("f2logs").Collection("logs")
	err := collection.Drop(ctx)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	return nil

}

// Update:
func (l *LogEntry) Update() (*mongo.UpdateResult, error) {
	// Declair variable that return from updated
	//var logUpdate LogEntry

	// create context if can't connect to database within 15 seconds then cancel
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// select databse and collection to work with
	collection := client.Database("f2logs").Collection("logs")

	// get log id with primitive objectid
	logId, err := primitive.ObjectIDFromHex(l.ID)
	if err != nil {
		fmt.Println("error: ", err)
	}

	// update
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": logId},
		bson.D{
			{"$set", bson.D{
				{
					"name",
					l.Name,
				},
				{
					"data",
					l.Data,
				},
				{
					"updated_at",
					l.UpdatedAt,
				},
			}},
		},
	)

	if err != nil {
		fmt.Println("error: ", err)
		return nil, err
	}

	return result, nil
}
