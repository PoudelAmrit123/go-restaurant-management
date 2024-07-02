package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "time"
)

func DBinstance() *mongo.Client {
	MongoDB := "mongodb://localhost:27017"

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MongoDB).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to mongoDB!!! ")
	}

	return client

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// defer cancel()

}

var Clinet *mongo.Client = DBinstance()

func OpenCollection(clinet *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = clinet.Database("restaurant").Collection(collectionName)

	return collection

}
