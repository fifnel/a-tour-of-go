package main

import "fmt"

// Struct Literals
// http://go-tour-jp.appspot.com/#28

type Vertex struct {
	X, Y int
}

var (
	p = Vertex{1, 2}
	q = &Vertex{1, 2}
	r = Vertex{X: 1}
	s = Vertex{}
)

func main() {
	fmt.Println(p, q, r, s)
}
