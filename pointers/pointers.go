package main

import (
	"fmt"
)

func main() {
	// when declaring a variable go will allocate memory for it in RAM
	// it would allocate memory for it depending on it's type
	// the memory would have a memory address which would go get the value when it is asked for it
	// the memory address is represented as a hexadecimal value
	a := 1

	// to access the address of a variable we use the ampersand followed by the variable 
	// when we assign a memory address of a variable to another variable using implicit initialization (:=)
	// the type on the left side is a pointer e.g var pointer *int`
	addr := &a
	fmt.Println("memmoery address of a: ", addr) // 0xc00001a100

	// dereferencing the pointer 
	// the * operator is called the dereferencing operator 
	// it returns the value stored in a particular memory address
	value := *addr
	fmt.Println("value stored in memory address 'addr':", value)

	// changing value using pointers 
	// we can change the value of a variable by reassigning a new value to it's memory address using pointers
	*addr = 2
	fmt.Println("new value of a:", a)

	// using functions 
	// to use the function we declared below whic expects a pointer
	// we pass in a pointer to the memory address using the & sign 
	fmt.Println("before change function: ", a)
	changeValue(&a)
	fmt.Println("after change function: ", a)


	// reassigning a value to a variable
	x := 5
	fmt.Println(&x)

	//
	y := x

	// you can also change the value of a variable by directly assigning a new value to it 
	x = 10
	// it still retains it's original memory address
	fmt.Println(&x)

	// we also see here that the value of y is unchanged because go assigns variables by ? 
	fmt.Println("y:", y, "x:", x)

}

// working with functions 
// when we want a function or method to mutate an external variable we use pointers 
// we can define the function to take in pointers 
// this function takes in a pointer of int type (a pointer that points to where an int value is stored)
func changeValue(x *int) {
	// change the value stored in the address x points to, to 15
	*x = 15
}