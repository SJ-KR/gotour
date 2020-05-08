package main

import (
	"fmt"
	"math"
)
// example 20
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// example 21
func pow(x, n ,lim float64) float64 {
	if v:=math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
// example 22
func pow2(x, n ,lim float64) float64 {
	if v:=math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}
func main() {
	//example 20
	fmt.Println("math.Sqrt() 2, -4 : ", math.Sqrt(2), math.Sqrt(-4))
	fmt.Println("custom sqrt() 2, -4 : ", sqrt(2), sqrt(-4))
	
	//example 21
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
	
	//example 22
	fmt.Println(pow2(3, 2, 10))
	fmt.Println(pow2(3, 3, 20))


}
