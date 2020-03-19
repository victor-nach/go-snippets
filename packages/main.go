package main

import (
	"fmt"
	"packages/one"
)

// when you try to import a package the go runtime would first go and look for that package in the GOROOT then next in the GOPATH
// when you are not working from the GOPATH, you need to use go modules
// the name that you referenced in the go mod init, is what you would use when referencing your packages internally

// youc can't have empty go files in an imported package, they must at least have a package declaration

// NESTED PACKAGES
// if a package has a nested package (folder), importing the parent would not automatically import the child package
// you also can't have nested folders with the same package name as the parent

// VARIABLE SCOPING
// consider all files in a package to be part of one big file that contains all of them
// a variable declared in the outer scope of one file cannot be re-declared in the outer scope of another file in the same package
// for variables declared in the outer scope of a package, they can be used before they are defined (also outer scope)
// they are all declared during initialization of the package

// Order of execution
// go run *.go
// ├── Main package is executed
// ├── All imported packages are initialized
// |  ├── All imported packages are initialized (recursive definition)
// |  ├── All global variables are initialized 
// |  └── init functions are called in lexical file name order
// └── Main package is initialized
//    ├── All global variables are initialized
//    └── init functions are called in lexical file name order

func main() {
	fmt.Println("main.go: welcome....")


	a := one.Person{
		Name: "victor",
	}

	fmt.Println("person", a)

	fmt.Println(one.Three)
}
