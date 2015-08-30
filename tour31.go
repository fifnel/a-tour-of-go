package main

import "fmt"

// Slicing slices
// http://go-tour-jp.appspot.com/#31

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])

	fmt.Println("p[:3] ==", p[:3])

	fmt.Println("p[4:] ==", p[4:])

	q := p[3:]
	fmt.Println("q ==", q)
	q[0] = 123
	fmt.Println("p ==", p)
	fmt.Println("q ==", q)
}
