package main

import (
	"fmt"
	"time"
)

// Switch evaluation order
// http://go-tour-jp.appspot.com/#46

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("tomorrow")
	case today + 2:
		fmt.Println("in two days")
	default:
		fmt.Println("too far away")
	}
}
