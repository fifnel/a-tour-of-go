package main

import "fmt"

// For is Go's "while"
// http://go-tour-jp.appspot.com/#19

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
