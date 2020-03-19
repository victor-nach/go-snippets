package one 

import (
	"packages/one/three"
)
// Unlike in languages like javascript, where you explicitly have to export things
// in Go variables and types, functions in the outer scope of the package that start with uppercase are automatically part of what is exported

type thing struct {
	name string
	color string
}


type Person struct {
	Name string
	Age string
}

var Three = three.Three