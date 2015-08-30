package main

import "fmt"

// Slices
// http://go-tour-jp.appspot.com/#30

func main() {
	//var p = []int{2, 3, 5, 7, 11, 13}
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] ==  %d\n", i, p[i])
	}
}
