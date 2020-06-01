package main

import (
	"fmt"
)

//These variables are package scoped.
var x int
var y string
var z bool

// The compiler will print the zeroes values of each identifier.
func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
