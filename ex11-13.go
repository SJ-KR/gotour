package main

import (
	"fmt"
)

// example 11
var x, y, z int
var c, python, java bool

// example 12
var x2, y2, z2 int = 1, 2, 3
var c2, python2, java2 = true, false, "no"

func main() {
	// example 11
	fmt.Println(x, y, z, c, python, java)
	// example 12
	fmt.Println(x2, y2, z2, c2, python2, java2)
	// example 13
	x3, y3, z3 := 5, 6, 7.7
	c3, python3, java3 := false, true, "yes"
	fmt.Println(x3, y3, z3, c3, python3, java3)
}
