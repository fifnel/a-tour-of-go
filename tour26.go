package main

import "fmt"

// Struct Fields
// http://go-tour-jp.appspot.com/#26

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
