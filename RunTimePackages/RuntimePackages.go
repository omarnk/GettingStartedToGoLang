package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 4
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Printf("%#U", x)
}
