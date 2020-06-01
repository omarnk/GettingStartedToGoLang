package main

import "fmt"

//Defination of new type which is newInt under UNDERLYING type int
type newInt int

//Variable creation of the new type.
var x newInt

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
