package main

import (
	"fmt"
)


func main() {
	// SLICE OF NUMBERS

	// a slice is declared using empty square brackets followed by the type followed by curly braces containing the elements
	nums := []int{1,2,3,4,5}
	fmt.Println("1.", nums)

	// the append function is a multipurpose function for working with slices, the first parameter is the slice that we want to work with 
	// append returns a new slice, which we are to store in place of the former slice

	// adding a new element to the slice
	// just add the slice, the the next argument which can one or more should be the elements we want to add
	nums = append(nums, 2) // adding one element
	fmt.Println("2.", nums)

	nums = append(nums, 4,13,7) // adding more than one element
	fmt.Println("3.", nums)

	//adding two slices
	nums2 := []int{21,22,25}
	nums = append(nums, nums2...) // appending another slice
	fmt.Println("4.", nums)


	// Trimming a slice
	// to trim a slice you just have to pass the start and stop index 
	// it would start at the start index and stop before the stop index
	// the original slice remains unchanged
	nums3 := append(nums[1:2]) // return elements from index 1 and stop before index 2
	fmt.Println("5.", nums3)
	nums3 = append(nums[:3]) // start index of 0, stop index of 3
	fmt.Println("6.", nums3)
	nums3 = append(nums[3:]) // start index should be 3, end index should be end of slice 
	fmt.Println("7.", nums3)

	// Removing elements from a slice
	nums = []int{1,2,3,4,5}
	// basically we create two sub-slices and they don't contain the elements we want to remove 
	nums3 = append(nums[:3], nums[4:]...) // remove the 3rd element
	fmt.Println("8.", nums3)

	// updating an element in a slice
	nums[3] = 18
	fmt.Println("9.", nums)
 


	// SLICE OF STRUCTS
	type Person struct {
		ID int
		Name string
		Age int
	}

	people := []Person{
		{
			ID: 1,
			Name: "victor",
			Age: 15,
		},
		{
			ID: 2,
			Name: "emmanuel",
			Age: 25,
		},
		{
			ID: 3,
			Name: "nonso",
			Age: 35,
		},
		{
			ID: 4,
			Name: "chukwudi",
			Age: 40,
		},
		{
			ID: 5,
			Name: "madubem",
			Age: 40,
		},
		{
			ID: 6,
			Name: "frank",
			Age: 40,
		},
		{
			ID: 7,
			Name: "john",
			Age: 40,
		},
	}

	// trim the slice 
	people2 := append(people[:2])
	fmt.Println("10", people2)

	// delete an element in the slice
	people2 = append(people[:3], people[4:]...) // remove the element at index of 3
	fmt.Println("11.", people2)

	// update an element in the slice 
	people[3] = Person{ // replaces the struct at that position, undefined fields take a null value - 0,"",false....
		Name: "anthony",
		Age: 25,
	}
	fmt.Println("12.", people)

	// update a value in a specific struct
	people[4].Name = "mercy"
	fmt.Println("13.", people)

	// get keys and values usiing reflect ??

	
	// filter slice
	// 	n := 0
	// for _, x := range a {
	// 	if keep(x) {
	// 		a[n] = x
	// 		n++
	// 	}
	// }
	// a = a[:n]
}