package main

import (
	"fmt"
)

func main() {
	// example 16
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum 1 to 9 : ", sum)
	// example 17, 18
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("powers of 2 closest to 1000 :", sum)
	// example 19
	/*
	for {
		//infinite loop
	}
	*/
}
