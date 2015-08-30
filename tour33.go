package main

import "fmt"

// Nil slices
// http://go-tour-jp.appspot.com/#33

// sliceの初期値はnil

func main() {
	var z []int
	printSlice(z)

	z = append(z, 123)
	z = append(z, 456)
	printSlice(z)

	z[2] = 123
	printSlice(z)
}

func printSlice(z []int) {
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}
