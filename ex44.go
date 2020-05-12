package main

import (
	"fmt"
)
func fib() func() int {
	a := 0
	b := 1
	return func() int {
		if a < b {
			a += b
			return b
		} else {
			b += a
			return a
		}
	}
}
func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
