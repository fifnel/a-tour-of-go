package main

import "fmt"

// Functions
// http://go-tour-jp.appspot.com/#7

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(123, 456))
}
