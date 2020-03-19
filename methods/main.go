package main

import (
	"fmt"
)

// methods are just functions with a receiver
// you can either use a pointer as the receiver or use a regular struct value
// if you use a value when that method is called a new copy is made and passed to the method
// when you use a pointer as the receiver, the underlying value can then be mutated
// a method must also be declared in the same package the receiving type was declared

type person struct {
	firstName 	string
	lastName 	string
	age			int
}

func (p *person) fullName() string {
	return p.firstName + " " + p.lastName
}


// you can write a method on non-struct types as well, just not on the type itself
// we can write a method on Myfloat but not on float64
type MyFloat float64

func (f MyFloat) Abs() {
}


func main() {
	a := person{
		firstName: "victor",
		lastName: "nonso",
	}

	// the fullName method requires a pointer receiver from its definition
	// but in Go it can also be called with a regular value of the same type
	// this also applies in the reverese when using a value receiver
	fmt.Println(a.fullName())
}