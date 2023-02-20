package dao

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	UsMongo *mongo.Collection

	client *mongo.Client
)

// connect to mongoDB
func connectMongoDB() {
	ctx := context.TODO()
	// connect to mongoDB
	fmt.Println("Connecting to MongoDB...")
	url := "mongodb://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PASSWORD") + "@" +
		os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT") + "/"

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}
	// get UserSession collection

	UsMongo = client.Database(os.Getenv("MONGO_DB")).Collection("user-session")

	fmt.Println("Connected to MongoDB")
}
