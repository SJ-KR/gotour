package main

import (
	"fmt"
)

func main() {
	p := []int {2, 3, 4, 5, 6, 7}

	//example 30
	fmt.Println("p==",p)
	for i:= 0; i<len(p); i++ {
		fmt.Printf("p[%d] : %d\n", i, p[i])
	}

	//example 31
	fmt.Println("p[1:4] =", p[1:4])

	fmt.Println("p[:3] =", p[:3])

	fmt.Println("p[4:] =",p[4:])

	fmt.Println("p[5:5] =",p[5:5])
}
