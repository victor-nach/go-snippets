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

	// makes a connection to the database and returns a new *mongo.client
	// takes in a context and an options parameter
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
	// if the database or collection does not exist, a new one is created
	// as usual, it has no structure so we can put anything inside 
	collection := client.Database("test").Collection("users")
	// collection

	// Create a single book
	user := User{
		FirstName: "Iheanacho Victor",
		LastName: "Nonso",
		Age: 24,
	}
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err)
	}
	// using the Insert one, mongo db only returns the object id back to you
	fmt.Println("successfully added", result.InsertedID)

	// INSERT MANY
	// to insert many documents we pass a slice
	users := []interface{}{
		User{
			FirstName: "daniel",
			LastName: "Anudu",
			Age: 15,
		},
		User{
			FirstName: "Anthony",
			LastName: "fred",
			Age: 76,
		},
		User{
			FirstName: "Amaka",
			LastName: "mgbaegbu",
			Age: 45,
		},
	}

	// returns an array of ids
	insertManyResult, err := collection.InsertMany(context.TODO(), users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(insertManyResult)
}