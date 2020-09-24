package main

import "fmt"

//Defination of new type which is newInt under UNDERLYING type int
type newInt int

//Variable creation of the new type.
var x newInt
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	//The conversion concept mentioned here, x has been converted to type int and the value,
	//has been assigned to identifier y.
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
