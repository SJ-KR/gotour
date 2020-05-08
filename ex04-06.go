package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Happy", math.Pi, "Day")
	//fmt.Println(math.pi)
	fmt.Printf("Now you have %g problems.\n", math.Nextafter(2,0))
}
