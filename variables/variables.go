package main

import (
	"fmt"
	"strconv"
)

// package level variables are defined outside the scope of any function in a go file
// they can be exported to other packages if the variable names are uppercased
// other files in this same package can have access to ALL package level variables (upper or lower case)
var b int = 65
// you can't do type inference for a package level variable
// a := 9
var BohemianAlz string = "dadan"
// also package level variables may or may not be used

// you can group variable declarations using parenthesis
var (
	x string = "fino alla fine"
	y int = 8
)


func main() {
	fmt.Println(b)

	//CONVERSION

	// There are different methods to help you convert types, usually the type name followed by parenthesis
	i := 4
	var j string 

	// trying to use the string() method to convert the number 4 would not work
	// it would look for what unicode character is stored at the the number 4, and that would be a comma ','
	j = string(i)
	fmt.Printf("%v, %T\n", j, j)

	// also a string is an immutable array of bytes

	// to properly convert back and forth between strings and numbers we can use the strconv library which comes with go

	// Itoa, meand Integer To Ask-key type, this method helps us convert from an integer to a string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)

	// but if you are converting between numeric types, you can just use the types as a function
	// int(), float32()...

	// Also every time you initialize a variable in Go, it has a zero value


	// BOOLEAN DATA
	var rand bool = true
	fmt.Printf("%v, %T\n", rand, rand)
	// the zero value for the boolean type is false 



	integers()
	strings()
}

func integers() {

	// INT TYPE (signed)
	// we have different kinds of integer types, ranging from the int8, int16, int32 and the int64 
	// the int8 is the smallest which is from -128 to 127

	// unsigned integer - uint8, uint16, uint32, but we don't have a 64 bit unsigned integer
	// unsigned integers mean that the integers do not have
	var n uint16 = 42
	fmt.Printf("%v, %T\n", n, n)

	// byte -  we also have byte which is an unsigned integer, a byte is an alias for an unsigned integer
	// because a lot data streams use this type to encode their data

	// during arithmetic operations, the type cannot change, also you cannot do arithmetic conversion with different types
	var aa int = 10
	var ab int8 = 3
	// aa + ab (would result in an error, because go would not do an implicit type conversion)
	// to use them together we have to do a type conversion
	fmt.Println(aa + int(ab))

	// also implicit type inference works to only certain types
	// for example you can't have a type implicitly defined as a uint8 or int64, or a float32, it takes the default type,e.g float64

}

func strings() {
	// STRINGS
	// string in Go are aliases for bytes
	// also note that strings are immutable
	// a string is treated as a collection of letters

	// declare a string
	s := "foodie" // [102 111 111 100]
	// convert it to a slice of bytes (byte array)
	b := []byte(s)
	// it converts each character to it's equivalent utf8 character or askey value
	fmt.Printf("%v, %T\n", b, b)

	
	// RUNE
	// we define string with double quotes and rune's with single quote
	// also a rune is just an alias for a int32 type (true type alias)
	var r rune = 'a'
	fmt.Printf("%v, %T", r, r)

	// We have a ReadRune method to read a rune
}

func constant() {

	// CONSTANTS

	// In other programming languages, constants are named using all uppercase variables
	// instead of the name MY_CONST, we use my_const as with other variables

	// to define a constant we use the const keyword 
	const a int = 14
	
	// constants in Go can be implicitly converted

	// Enumerated constants ??
}
