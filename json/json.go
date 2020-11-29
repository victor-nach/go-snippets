package main

import (
	"encoding/json"
	"fmt"
)

type House struct {
	Address 	string		`json:",omitempty"`
}

	// json tags
	// if the json key is not available, it uses the same name as the field (default)
	// if we use the "omitempty"  option in the field tag, the field would be ommited if that field is empty
	// empty values can be any of - defined as false, 0, a nil pointer, an empty array, slice, map or string
	// if we use the "-" option, then that field is always ommited
type Person struct {
	FName 		string 	`json:"firstName"`
	LName		string	`json:"lastName,omitempty"`
	Age			int		`json:"-"`
	Address		string
	nationality	string
	Houses		*House	`json:"houses,omitempty"`
}


type Target struct {
	Id			string	`json:"id"`
	Message 	string	`json:"message"`
}

	// Tags provide conventions for built-in parsing, they are not strict rules 
	// Different packages use tags for various reasons e.g json,bson,gorm .....

	
func main() {
	// Marshall returns the json encoding of a data interface as a byte array
	// the encoding of each field can be customized by the format string stored under the "json" key in the struct
	// emma := Person{
	// 	FName: "emmanuel",
	// 	LName: "Iheanacho",
	// 	Age: 23,
	// 	nationality: "Nigerian",
	// }
	byte, _ := json.Marshal(Person{
		FName: "emmanuel",
		LName: "Iheanacho",
		Age: 23,
		nationality: "Nigerian",
	})
	// fmt.Println(byte) // returns a lenghty byte array
	// fmt.Println(string(byte)) 

	//marshall for maps

	// unmarshall takes in a byte array containing the json encoding and also a pointer to a struct, map...
	// Unmarshall parses the json encoded data(byte array) and stores the result in the value pointed to by the data interface
	// it does the reverse of marshall, using the same rules set by tags
	var victor Person
	json.Unmarshal(byte, &victor)
	fmt.Println(victor)

	payload, _ := json.Marshal(Target{
		Id: "12345",
		Message: "hello from this side",
	})
	// fmt.Println(payload)
	var victorr Target
	json.Unmarshal(payload, &victorr)
	fmt.Println(victorr)
	
	// Encoding
	// An Encoder(type Encoder) writes json values to an output stream

	// NewEncoder 
	// takes in a stream(io.writer) and returns a new *Encoder that writes to w
	// essentially we pass in a writter like http.ResponseWriter, then this returns an Encoder
	// We can now use the Encode function to write a json value to that stream, it takes in a data interface

	// Encode - func(*Encode) Encode(v interface{}) error
	// Encode is a method available to the Encoding interface

	// Decoding
	// An Decoder(type Decoder) reads json values from an input stream


}