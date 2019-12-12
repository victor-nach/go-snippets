package main

import (
	"context"
	"fmt"
	"log"
	// "time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID			primitive.ObjectID		`json:"id,omitempty" bson:"_id,omitempty"` 
	// note that you don't have to specify the bson, automatically mongo adds an id field and also all the field names in a struct are changed to lower case in the collection
	FirstName 	string					`json:"firstName,omitempty" bson:"firstname,omitempty"`
	LastName  	string					`json:"lastName,omitempty" bson:"lastname,omitempty"`
	Age       	int						`json:"age,omitempty bson:"age,omitempty"`
}

type Book struct {
	ID 				primitive.ObjectID
	Title        string
	Author       string
	Type         string
	ReadersCount *string
}

var client *mongo.Client

func main() {
	// .Clients() creates a new clients options instance
	// .ApplyURI parses the application URI and then adds it to the client options instance
	clientOptions := options.Client().ApplyURI("mongodb+srv://victor:victor_123@victor-nach-gq3xj.mongodb.net/test?retryWrites=true&w=majority")

	// makes a connection to the database and returns a new *mongo.client
	// takes in a context and an options parameter
	var err error
	// ctx := context.WithTimeout(context.Background(), 10*time.Second) // alternative context to use
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// verifies that the client can connect to the topology
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB")

	// get a handle for a particular collection
	// collection := client.Database("test").Collection("users")

	Create()
	Retreive()
	Update()
}

func Create() {
	fmt.Println("\nCreate........")
	// get a handle for a particular collection
	// if the database or collection does not exist, a new one is created
	// as usual, it has no structure so we can put anything inside
	collection := client.Database("test").Collection("users")
	// collection

	// Create a single book struct
	user := User{
		FirstName: "Iheanacho Victor",
		LastName:  "Nonso",
		Age:       24,
	}

	// then we pass in this struct to the collection.insertOne with a context
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Println(err)
	}
	// using the Insert one, mongo db only returns the object id back to you
	fmt.Println("successfully added", result.InsertedID)

	// INSERT MANY
	// to insert many documents we pass a slice
	// users := []interface{}{
	// 	User{
	// 		FirstName: "daniel",
	// 		LastName:  "Anudu",
	// 		Age:       15,
	// 	},
	// 	User{
	// 		FirstName: "Anthony",
	// 		LastName:  "fred",
	// 		Age:       76,
	// 	},
	// 	User{
	// 		FirstName: "Amaka",
	// 		LastName:  "mgbaegbu",
	// 		Age:       45,
	// 	},
	// }

	// returns an array of ids
	// insertManyResult, err := collection.InsertMany(context.TODO(), users)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(insertManyResult)
}

func Retreive() {
	fmt.Println("\nRetreive........")
	// to retreive we need to use a filter

	// first we get a handle to the collection we want to insert into
	collection := client.Database("test").Collection("users")

	// FIND ONE USER
	var user User

	// when writing your filters, you have to use the fieldnames as stored in the collection or as defined in the bson tags on the structs
	// whuch would be different from the structs names(lowercase)
	// a bson document could contain other bson.E literals, in the case below, the inner object is converted to bson.E
	filter := bson.D{{"firstname", "daniel"}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	// the returned struct doesn't contain the bson id if not added to the struct's field
	log.Println("find one using bson: ", user)

	// err = collection.FindOne(context.TODO(), User{FirstName: "Anthony"}).Decode(&user) // didn't work when using structs

	// using multiple filter options we use bson.M, that contains an "$or" and an array of of other bson documents
	// "$or": []bson.M{ bson.M{}, bson.M{} }
	filter2 := bson.M{
		"$or": []bson.M{
			bson.M{"firstname": "Anthony"},
			bson.M{"lastname": "fred"},
		},
	}
	err = collection.FindOne(context.TODO(), filter2).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	log.Println("Using multiple filter options: ", user);
	

	// Searching using mongodb object id
	id, _ := primitive.ObjectIDFromHex("5ddfdfea5d248400d5be16b7")
	err = collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	log.Println("Using object id to search: ", user);

	// FIND ALL USERS

	var users []User
	// we use the options.FindOne() to set custom delimeter for the find
	cur, err := collection.Find(context.TODO(), bson.D{{}}, options.Find().SetLimit(5))
	if err != nil {
		log.Println(err)
	}

	// it is advisable to defer closing of the cursor immediately after you get it
	defer cur.Close(context.TODO())

	// check for errors in getting the cursor
	if err = cur.Err(); err != nil {
		log.Println(err)
	}
	// the cursor returned contains an loopable collection of all the data found
	for cur.Next(context.TODO()) {
		var user User
		cur.Decode(&user)
		users = append(users, user)
	}
	log.Println("all users in the db: ", users)
}

func Update() {
	log.Println("\n Update...........")

	collection := client.Database("test").Collection("users")
	// to update our documents, we need bson documents for both the filters and the updates
	// note that a bson map has to always have a key and value
	filter := bson.M{
		"$and": []bson.M{
				bson.M{"firstname": "daniel"},
				bson.M{"lastname": "Anudu"},
		},
	}

	query := bson.M{
		"$set": bson.M{
			"age": 22,
		},
	}
	
	// for the update one query, the updated result is returned
	var user User
	updateResult, err := collection.UpdateOne(context.TODO(), filter, query)
	if err != nil {
		log.Println(err)
	}
	err = updateResult.UnmarshalBSON(&user)
	log.Println(&updateResult)
}