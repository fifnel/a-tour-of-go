package main

import "fmt"

// Named results
// http://go-tour-jp.appspot.com/#10

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
