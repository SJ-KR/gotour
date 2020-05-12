package main

import (
	"fmt"
	"math/cmplx"
)
func Cbrt(x complex128) complex128 {
	var z complex128 = 1
	for i := 0; i < 10; i++ {
		z = z - (cmplx.Pow(z,3)-x)/(3*cmplx.Pow(z,2))
		//fmt.Println(z)
	}
	return z
}
func main() {
	fmt.Println(Cbrt(-8))
}
