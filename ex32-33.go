package main

import (
	"fmt"
)

func main() {
	//example 32
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	//example 33
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
	z = make([]int, 3, 5)
	z[0] = 10
	z[2] = 30
	fmt.Println(z, len(z), cap(z))
	z = z[:cap(z)]
	fmt.Println(z, len(z), cap(z))

}
//example 32
func printSlice(x string, y []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",x, len(y), cap(y), y)
}
