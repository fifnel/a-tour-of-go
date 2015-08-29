package main

import "fmt"

// Structs
// http://go-tour-jp.appspot.com/#25

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
