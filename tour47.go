package main

import (
	"fmt"
	"time"
)

// Switch with no condition
// http://go-tour-jp.appspot.com/#47

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("good morning!")
	case t.Hour() < 17:
		fmt.Println("good afternoon!")
	default:
		fmt.Println("good evening!")
	}
}
