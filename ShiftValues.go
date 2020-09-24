package main

import (
	"fmt"
)

func main() {
	x := 50
	var y = x << 1
	fmt.Printf("%b\t%d\t%#x\t\n", x, x, x)
	fmt.Printf("%b\t%d\t%#x\t", y, y, y)
}
