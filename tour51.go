package main

import (
	"fmt"
	"math"
)

// Methods continued
// http://go-tour-jp.appspot.com/#51

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
