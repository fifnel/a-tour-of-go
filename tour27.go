package main

import "fmt"

// Pointers
// http://go-tour-jp.appspot.com/#27

type Vertex struct {
	X int
	Y int
}

func main() {
	p := Vertex{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(p)
}
