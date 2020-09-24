package main

import (
	"fmt"
)

func main() {
	a := (50 == 50)
	b := (50 <= 60)
	c := (70 >= 80)
	d := (80 != 70)
	e := (80 < 70)
	f := (70 > 100)
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}
