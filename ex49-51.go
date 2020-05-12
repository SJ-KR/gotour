package main

import (
	"fmt"
	"math"
)
// example 50
type Vertex struct {
	X, Y float64
}
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
// example 51
type Myfloat float64

func (f Myfloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	// example 50
	v := &Vertex{3,4}
	fmt.Println(v.Abs())

	// example 51
	f := Myfloat(-math.Sqrt2)
	fmt.Println(f.Abs2())
}
