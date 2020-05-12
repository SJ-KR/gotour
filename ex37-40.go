package main

import (
	"fmt"
)
type Vertex struct {
	Lat, Long float64
}
// example 37
var m map[string]Vertex

// example 38
var m2 = map[string]Vertex {
	"Bell": Vertex {
		11.111, 22.222,
	},
	"Google": Vertex {
		33.333, 44.444,
	},
}
// example 39
var m3 = map[string]Vertex {
	"Kakao" : {555.55, 666.66},
	"Samsung": {777.77, 888.88},
}

func main() {
	// example 37
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex {
		40.123, -70.456,
	}
	fmt.Println(m["Bell Labs"])

	// example 38
	fmt.Println(m2)

	// example 39
	fmt.Println(m3)

	// example 40
	m4 := make(map[string]int)

	m4["Answer"] = 42
	fmt.Println(m4["Answer"])
	m4["Answer"] = 48
	fmt.Println(m4["Answer"])

	v, ok := m4["Answer"]
	fmt.Println("The value :",v,"Present? :",ok)

	delete(m4,"Answer")
	fmt.Println(m4["Answer"])

	v, ok = m4["Answer"]
	fmt.Println("The value :",v,"Present? :",ok)
}
