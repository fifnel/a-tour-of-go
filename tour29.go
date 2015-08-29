package main

import "fmt"

// The new function
// http://go-tour-jp.appspot.com/#29

type Vertex struct {
	X, Y int
}

func main() {
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)

	var vv *Vertex
	vv = &Vertex{123, 456}
	fmt.Println(vv)
}
