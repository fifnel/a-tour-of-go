package main

// Multiple results
// http://go-tour-jp.appspot.com/#9

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := "hello", "world"
	fmt.Println(a, b)
	a, b = swap("hello", "world")
	fmt.Println(a, b)
}
