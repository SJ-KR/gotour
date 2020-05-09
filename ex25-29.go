package main

import (
	"fmt"
)
type Vertex struct {
	X int
	Y int
}
// example 28
var (
	a = Vertex{1,2}
	b = &Vertex{3,4}
	c = Vertex{X:10}
	d = Vertex{}
)
func main() {
	//example 25
	fmt.Println(Vertex{1, 2})

	//example 26
	v := Vertex{3, 4}
	v.X = 6
	fmt.Println(v)

	//example 27
	p := Vertex{10, 12}
	q := &p
	r := p
	q.X = 1e9
	fmt.Println("p : ",p)
	fmt.Println("q : ",q)
	fmt.Println("r : ",r)
	
	//example 28
	e := *b
	fmt.Println(a, b, c, d, e)
	e.X = 20
	fmt.Println(a, b, c, d, e)
	b.X = 30
	fmt.Println(a, b, c, d, e)

	//example 29
	t := new(Vertex) // var t *Vertex = new(Vertex)
	fmt.Println(t)
	t.X = 11
	fmt.Println(t)

}
