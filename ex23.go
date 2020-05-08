package main

import (
	"fmt"
	"math"
)
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x) / (2*z)
		//fmt.Println(z)
	}
	return z
}
func main() {
	a := 2.0
	fmt.Println("Newton's method root",a," : ",Sqrt(a))
	fmt.Println("math.Sqrt(",a,") : ",math.Sqrt(a))

}
