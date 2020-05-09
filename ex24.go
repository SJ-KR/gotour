package main

import (
	"fmt"
	"math/cmplx"
)
var (
	ToBe bool = false
	s string = "hello"
	Maxint64 uint64 = 1<<64 - 1
	Maxint32 uint32 = 1<<32 - 1
	Maxint16 uint16 = 1<<16 - 1
	z complex128 = cmplx.Sqrt(-5+12i)
)
func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, s, s)
	fmt.Printf(f, Maxint64, Maxint64)
	fmt.Printf(f, Maxint32, Maxint32)
	fmt.Printf(f, Maxint16, Maxint16)
	fmt.Printf(f, z, z)
}
