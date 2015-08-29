package main

import "fmt"

// For continued
// http://go-tour-jp.appspot.com/#17

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
