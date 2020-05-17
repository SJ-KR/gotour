package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Print(<-ch," ")
	}
	fmt.Println()

	fmt.Println("New(1), New(1) : ", Same(tree.New(1), tree.New(1)))
	fmt.Println("New(1), New(2) : ", Same(tree.New(1), tree.New(2)))
	fmt.Println("New(2), New(2) : ", Same(tree.New(2), tree.New(2)))

}
