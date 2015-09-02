package main

import (
	"fmt"
	"time"
)

// Errors
// http://go-tour-jp.appspot.com/#55

type MyError struct {
	When time.Time
	What string
}

// errorインターフェースの Error() string の実装
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	me := &MyError{
		time.Now(),
		"it didn't work",
	}

	return me
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
