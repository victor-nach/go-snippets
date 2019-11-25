package main

import (
	"fmt"
)

func main() {

	array()
	slice()
	stactOps()
}

func array() {
	// Array is a collection of items with the same size
	// to define an array, we need to specify the length and type of the array
	// Arrays have a fixed size, whose value needs to be known at compile time
	grades := [3]int{20, 23, 98}
	fmt.Printf("grades:  %v\n", grades)

	// We can also define to define the array without defining the exact number by using three dots '...'
	a := [...]int{12,23,14,73}
	fmt.Printf("length of the array: %v\n", len(a))
	// we can use the len method to get the length of an array 

	// You can use the double braces to define multi-dimensional arrays [2][3]int
	multiArr := [3][3]int{ [3]int{1, 4, 6}, [3]int{23, 1, 7}, [3]int{9, 9, 4}  }
	fmt.Printf("multi-dimensional array: %v\n", multiArr)

	// When we re-assign an array to another value, the values are copied except we use a pointer
	b := a
	// b would have it's own values different from a
	fmt.Println(b)

}

func slice() {
	// For slices, we do not need to define the initial size
	a := []int{1,2,3,4,5,6,7,8,9,10}

	// Every slice has a backing array, 
	// and since the lengths of slices are not specified, to get the capacity of the backing array we use the cap() method
	fmt.Printf("Capacity: %v\n", cap(a))
	// in this case the underlying array is the same as the underlying array

	// Also slices are passed by references, meaning unlike arrays, slices aren't copied
	// So a change to the value of one of the referencies, would affect the other one


	// There are other ways to create slices without using the literal syntax

	b := a[:] //slice all the elements
	fmt.Println("b:", b)

	c := a[3:] // starts from the element at index 3 to the end
	fmt.Println("c: ", c)

	d := a[:6] // starts from the beginning and ends before the element at index 6
	fmt.Println(d)

	e := a[3:6] // starts from the element at index 3 to the element before the index 6
	fmt.Println(e)


	// MAKE
	// we can also use the make function to create a slice
	// the make function for slices, takes in the type, the length and the capacity
	x := make([]int, 3, 100)
	// when we do this x is defined to it's zero value
	fmt.Println("x: ", x) // [0,0,0]
	fmt.Printf("length: %v\n", len(x)) // prints 3
	fmt.Printf("capacity: %v\n", cap(x)) // prints 100


	j := []int{} // an empty slice with length and capacity = 0
	fmt.Println("j: ", j)
	fmt.Printf("length of j: %v\n", len(j)) 
	fmt.Printf("capacity of j: %v\n", cap(j))

	// A slice's size can change over time, elements can be added or removed from the slice using the append function
	// The append function takes two or more elements, the first one being the source slice
	// if the capacity of the underlying array, append returns a new slice after copying the values of the previous slice
	// it creates an underlying array big enough to hold the slice
	
	j = append(j, 1)
	fmt.Println("j: ", j)
	fmt.Printf("length of j: %v\n", len(j)) 
	fmt.Printf("capacity of j: %v\n", cap(j))

	j = append(j, 2, 3, 4)
	fmt.Println("j: ", j)
	fmt.Printf("length of j: %v\n", len(j)) 
	fmt.Printf("capacity of j: %v\n", cap(j))

	j = append(j, 5)
	fmt.Println("j: ", j)
	fmt.Printf("length of j: %v\n", len(j)) 
	fmt.Printf("capacity of j: %v\n", cap(j)) // the capacity is increased to 8, which is two times the former capacity (4) even tho, the length is just 5

	// In each of this case a new slice is always returned, if you can determine the capacity of a slice, it's better to define it at runtime

	// you can also spread a slice inside an append function, because the append function only takes in one or more individual parameters
	k := []int{12,4,7}
	j = append(j, k...)
	fmt.Println(j)
}

func stactOps() {

	a := []int{1,2,3,4,5}

	fmt.Println("\nUnshifting")
	// Unshift
	// Remove the first element in the slice
	b := a[1:] // returns the elements after the first element in slice a
	fmt.Println(b)

	// pop
	// remove the last element
	c := a[:len(a) -1]
	fmt.Println(c)

	// Remove elements
	d := append(a[:2], a[3:]...)
	fmt.Println(d)

	// One thing to note is that when doing things like this, we make changes to the underlying array
	// the only way to secure from this is to create a loop in which we first make a copy of the original array

}