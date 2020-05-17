package main

import (
	"fmt"
)
// example 64
func sum (a []int, c chan int) {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	c <- sum
}
// example 66
func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
// example 67
func fib2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
		}
	}
}
func main() {
	// example 64
	fmt.Println("-----example 64-----")
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	// example 65
	fmt.Println("-----example 65-----")
	c2 := make(chan int, 5)
	c2 <- 1
	c2 <- 2
	fmt.Println(<-c2)
	fmt.Println(<-c2)

	// example 66
	fmt.Println("-----example 66-----")
	c3 := make(chan int, 10)
	go fib(cap(c3), c3)
	for i := range c3 {
		fmt.Println(i)
	}

	// example 67
	fmt.Println("-----example 67-----")
	c4 := make(chan int)
	quit := make(chan int)
	go func() {
		for i:=0; i<10; i++ {
			fmt.Println(<-c4)
		}
		quit <-0
	}()
	fib2(c4,quit)
}
