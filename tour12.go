package main

import "fmt"

// Variables with initializers
// http://go-tour-jp.appspot.com/#12

var i, j int = 1, 2
var c, python, java = true, false, "no!"

func main() {
	fmt.Println(i, j, c, python, java)
}
