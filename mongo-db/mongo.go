package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
}

type Book struct {
	Title   		string
	Author  		string
	Type    		string 
	ReadersCount   	*string
}

var client *mongo.Client

func main() {
	// .Clients() creates a new clients options instance
	// .ApplyURI parses the application URI and then adds it to the client options instance
	clientOptions := options.Client().ApplyURI("mongodb+srv://victor:victor_123@victor-nach-gq3xj.mongodb.net/test?retryWrites=true&w=majority")

	// returns a new *mongo.client
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// verifies that the client can connect to the topology
	err = client.Ping(context.TODO(), nil)

	if err!= nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB")

	// get a handle for a particular collection
	// collection := client.Database("test").Collection("users")

	Create()
}

func Create() {
	// get a handle for a particular collection
	collection := client.Database("test").Collection("users")
	collection
}