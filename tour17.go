package main

import "fmt"

// For
// http://go-tour-jp.appspot.com/#17

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
