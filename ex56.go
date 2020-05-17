package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number : %f\n",float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f >= 0 {
		return f, nil
	}
	return -1, ErrNegativeSqrt(f)
}

func main() {
	if r, e := Sqrt(2); e != nil {
		fmt.Println(r, e)
	}
	if r, e := Sqrt(-2); e != nil {
		fmt.Println(r, e)
	}
	fmt.Println(Sqrt(10))
	fmt.Println(Sqrt(-10))
}
