package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Age       int
}

type Book struct {
	gorm.Model
	Title  	string
	Author 	string
	// using default option in the gorm tag tells gorm to use that value whenever a zero value is assigned to it on creation
	Type	string `gorm:"default:'hard-cover'"`
	// zero values include 0, "", false..., if in the case that is a possible value we can use a pointer as the type instead
	// so if you pass zero it would actually take in zero and not change it to 1
	ReadersCount	*string `gorm:"default:1"`
}

func main() {
	connectionString := "postgres://ylhighgx:LnyPyXA7noS7dXybLNIN0cSwVLyxotBi@isilo.db.elephantsql.com:5432/ylhighgx"

	var err error
	db, err = gorm.Open("postgres", connectionString)

	if err != nil {
		log.Println("Could not connect to the db")
		log.Fatal(err)
	}
	log.Println("Succesfully connected to the db")
	// close connection after the call to the database connection
	defer db.Close()

	// to migrate, we pass in an empty struct, gorm creates a new table based on the fields in that struct and gorm tags
	// auto migrate missing fields, but would not delete/change current data

	// uncomment when changes have been made to the user struct or when starting fresh
	// db.AutoMigrate(&User{})
	log.Println("Migrated user table")

	// // uncomment when changes have been made or when starting fresh
	// db.AutoMigrate(&Book{})
	log.Println("Migrated book table")

	// call the create methods
	createBook()
	createUser()
	
}

func createUser() {
	newUser := User{
		FirstName: "Victor",
		LastName:  "Iheanacho",
		Age:       23,
	}

	// db.create takes in a memory address and it also updates the value stored in that address after creating the field in the db
	// it also automatically knows which table to create an entry in based on the type used
	db.Create(&newUser)
	// after creation of the user, the user struct we passed to it now has an ID field gotten from the auto-increment in the db table
	fmt.Println(newUser.ID)

	// we can also add the .Error option to check for any errors, this is always advisable 
	err := db.Create(&newUser).Error
	if err != nil {
		log.Fatal(err)
	}


	// the create method used above returns a db that points to the user table
	// doing this would attempt to create a new entry in that same table as well
	// newBook := Book{
	// 	Title: "little fears",
	// 	Author: "james wilson",
	// }
	// err = db.Create(&newBook).Error
	// if err != nil {
	// 	log.Println(err)
	// }
}

func createBook() {
	newBook := Book{
		Title: "little fears",
		Author: "james wilson",
	}

	// the create method used above returns a db that points to the user table
	// doing this would attempt to create a new entry in that same table as well
	err := db.Create(&newBook).Error
	if err != nil {
		log.Println(err)
	}
	fmt.Println(newBook.ID)
}
