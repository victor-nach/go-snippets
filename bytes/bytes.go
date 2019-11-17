package main

import "fmt"

func main() {
	// byte is an alias for uint8 to distinguish from 8-bit unsinged integers 
	// uint8 is the set of all unsigned 8-bit integers ranfing from 0-255
	// here we are casting the string "random" to a byte array
	a := []byte("fund")
	fmt.Println(a); // [102 117 110 100]

	// strings are immutable whereas bytes are not
	// we can add a letter to the string "fund after converting it to bytes"
	// there are so many other methods that take in bytes eg. split, replaceAll, contains....
	a = append(a, byte('s'))
	fmt.Println(a) // [102 117 110 100 115] 

	// we can convert a back to a string to see the output
	fmt.Println(string(a)) // funds

}