package main

import (
	"fmt"
)

// example 7, 8
//func add(x int, y int) int {
func add(x, y int) int {
	return x + y
}

// example 9
func swap(x, y string) (string, string) {
	return y, x
}

// example 10
func split(sum int) (x, y int) {
	x = sum / 10
	y = sum % 10
	return
}
func main() {
	// example 7, 8
	fmt.Println(add(42, 13))
	// example 9
	a, b := "hello", "world"
	fmt.Println(swap(a, b))
	// example 10
	fmt.Println(split(78))
}
