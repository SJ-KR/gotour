package main

import (
	"fmt"
	"math"
)
type Abser interface {
	Abs() float64
	Pow() float64
}
func main() {
	var a Abser
	f := Myfloat(-math.Sqrt2)
	v := Vertex{3,4}

	a = f
	fmt.Println(a.Abs(), a.Pow())

	a = &v
	//a = v

	fmt.Println(a.Abs(), a.Pow())
}
type Myfloat float64
func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func (f Myfloat) Pow() float64 {
	return float64(f*f)
}
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (v *Vertex) Pow() float64 {
	return float64(v.X*v.X + v.Y*v.Y)
}
