package main

import (
	"fmt"
)

//These variables are package scoped.
var x int = 42
var y string = "James Bond"
var z bool = true

// The compiler will print the zeroes values of each identifier.
func main() {
	//Identifier s will print the default values of the identifiers (x ,y and z)
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)

	//Identifier typeofVars will print the types of the Identifiers (x ,y and z)
	typeofVars := fmt.Sprintf("%T\t%T\t%T", x, y, z)
	fmt.Println(typeofVars)
}
