package main

import (
	"fmt"
	"time"
)
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(), "it didn't work",
	}
}
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
	var err2 error
	if err2 = run(); err2 != nil {
		fmt.Println(err2)
	}
}
