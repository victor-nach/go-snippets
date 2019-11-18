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
	// log.Println("Migrated user table")

	// // uncomment when changes have been made or when starting fresh
	// db.AutoMigrate(&Book{})
	// log.Println("Migrated book table")

	// call the create methods
	create()
	// Retreive()
	Update()
}

func create() {
	newUser := User{
		FirstName: "Victor",
		LastName:  "Iheanacho",
		Age:       23,
	}

	// db.create takes in a memory address and it also updates the value stored in that address after creating the field in the db
	// it also automatically knows which table to create an entry in based on the type used
	db.Create(&newUser)
	// after creation of the user, the user struct we passed to it now has an ID field gotten from the auto-increment in the db table
	fmt.Println("user id after query: ",newUser.ID)

	// we can also add the .Error option to check for any errors, this is always advisable 
	// err := db.Create(&newUser).Error 


	newBook := Book{
		Title: "little fears",
		Author: "james wilson",
	}
	// it automatically knows which table to put the data in
	err := db.Create(&newBook).Error
	if err != nil {
		log.Println(err)
	}
	fmt.Println("book id after query i: ", newBook.ID)
}

func Retreive() {
	// find all 
	//first we declare an array of structs to hold the data
	// then we pass a pointer to that address using gorm 

	// find all users 
	var users []User
	db.Find(&users) // SELECT * FROM users;
	// the users slice declared above would now contain all the users in the user table
	// fmt.Println(users)  

	// find all books 
	var books []Book
	db.Find(&books) // SELECT * FROM books;
	// fmt.Println(books[0])

	// find by ID 
	var user User
	// this pattern only works where the id is the primary key and is also an integer 
	db.First(&user, 10) // select all from users where id = 10

	// find by attribute 

	// plain method
	// getting just one 
	db.Where("firstname = ?", "emma").First(&user) // select * from users where firstname = emma limit 1
	// getting all that match where clause 
	db.Where("firstname = ?", "emma").Find(&users) // select * from users where firstname = emma

	// struct method 
	db.Where(&Book{Title: "john", Author: "cole" }).Find(&user) // SELECT * FROM users WHERE Title = "john" AND Author = "cole" LIMIT 1;

	//inline method
	// these methods take the output struct as the first value 
	db.Find(&users, User{FirstName: "alhaji", LastName: "rozy"}) //struct
	db.First(&user, "firstname = ?", "emma") // returns the first match



}

func Update() {
	// to update fields using gorm we can first get a reference to that user using:
	var user User

	// select * from users where firstname = emma limit 1
	// db.Where("firstname = ?", "emma").First(&user)
	// or 
	db.First(&user, 2) // select all from users where id = 2
	fmt.Println("selected user id", user.ID)

	// then we can now use the user we have as a reference for where we want to update 

	// METHOD 1:
	// we can change the values before calling save and passing an updated struct
	// we can change the values we want
	user.Age = 14
	db.Save(&user) //save would include all fields when updating the field, even if they were not changed
	// note that this is UPDATE users ... (it updates more than one user if matched)
	fmt.Println("updated user: ", user)

	// METHOD 2:
	// we can use updates with a struct containing updated fields, note that only changed fields and non-empty fields would be updated
	// we can also add more than one check using where
	// UPDATE users SET first_name = 'victor', age = 24, updated_at = '....' WHERE id = 2 AND age = 7
	// db.Model(&user) is used to specify the user we want to change 
	db.Model(&user).Where(&User{Age: 7}).Updates(User{FirstName: "musa", Age: 24})
	fmt.Println("updated user: ", user)

	// we can use .Select and .Omit to specify fields we want or do not want to be updated
	// db.Model(&user).Select("") // ** test to see if .Select works with struct
}

func Delete() {
	//we use the .Delete method to delete records as well
	// just as in the case for the update method we need to first select the user
	var user User
	db.First(&user, 2)

	// when using delete gorm uses the primary key to delete the record 
	// you need to make sure that a primary key exists if not gorm would delete all records !!! 
	db.Delete(&user)
	// Also .Delete automatically soft deletes if the field has a delete_at column, it doesnt actually remove it
	// it unscopes it so that when you try to pull records it won't show up

	// to view a field that has been soft deleted, you can use db.Unscoped().Where()...
	var users []User
	db.Unscoped().Where(&User{Age: 4}).Find(&users) // would return all values in and out of scope (deleted)

	// to delete a record permanently
	// db.Unscoped().Delete(&user) 

}
 func ErrorHandling() {
	 // how to handle errors
 }
