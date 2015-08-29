package main

import "fmt"

// Short variable declarations
// http://go-tour-jp.appspot.com/#13

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
