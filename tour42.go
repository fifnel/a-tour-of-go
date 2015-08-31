package main

import (
	"fmt"
	"math"
)

// Function values
// http://go-tour-jp.appspot.com/#42

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))
	fmt.Println(hypot)
}
