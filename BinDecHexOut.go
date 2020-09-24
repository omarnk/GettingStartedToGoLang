package main

import (
	"fmt"
)

var x int = 500
var y int = 90

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%d\t%b\t%#x\t", y, y, y)

}
