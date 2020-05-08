package main

import (
	"fmt"
)
// example 14
const Pi = 3.14
// example 15
const (
	Big = 1 << 100
	Small = Big >> 99
)
func needint(x int) int { return x*10 }
func needfloat(x float64) float64 {
	return x*0.1
}
func main() {
	// example 14
	const World = "안녕"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	// example 15
	fmt.Println("Small : ", Small)
	fmt.Println("Big : ", float64(Big))
	fmt.Println(needint(Small))
	fmt.Println(needfloat(Small))
	fmt.Println(needfloat(Big))
}
