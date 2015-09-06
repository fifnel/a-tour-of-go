package main

import (
	"fmt"
	"time"
)

// Goroutines
// http://go-tour-jp.appspot.com/#63

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("World")
	say("hello")
}
