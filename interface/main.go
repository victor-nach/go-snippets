package main 

import (
	"fmt"
)

// interfaces are used when we really care about the methods
// so and interface type is a set of methods
// any value of any underlying type that implements the methods specified in the interface, implements that interface
// there is no explicit declaration of intent when implementing an interface

// with interfaces you could write a function that needs to run a set of methods 
// and have that function take that interface as parameter, this way it can be used by multiple methods

// and interface carries two things, information about the type and the data in the interface
// the type information contains information that is used to convert back to the underlying type using type assertion
// data contains information that the interface holds, which comes from the underlying type's value (the value passed when creating the interface)
// if a function requests for an interface, you can create an interface type first from a different type that implements in then pass it in or
// pass in a type that implements that interface. if you pass in a type that implements that interface, go automatically converts it to the interface type first
// so to actually get access to the data, you need to convert it back to underlying type using type assertions i.(T)
// since an interface can be implemented by different types, you need to know the type you are converting is the correct one, if not a panic would occur
// that's why the type assertion returns a boolean to let you know if it is of the proper type or not eg. { value, ok := i.(T) }

// empty data interface
// when a function takes in an empty data interface

type interfaceOne interface {
	greet() string
}

type person struct {
	name string
}

func (p person) greet() string {
	return p.name + "says hi.."
}
func main() {
	a := person{"victor"}
	doOne(a)
	empty(a)
}

func doOne(i interfaceOne) {
	fmt.Println(i.greet())
	b := i.(person)
	fmt.Println(b.name)
}

func empty(i interface{}) {
	b := person(i)
	fmt.Println(b.name)
}