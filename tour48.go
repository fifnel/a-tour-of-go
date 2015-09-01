package main

import (
	"fmt"
	"math/cmplx"
)

// Advanced Exercise: Complex cube roots
// http://go-tour-jp.appspot.com/#48
// http://c4se.hatenablog.com/entry/2013/01/19/164359

func Cbrt(x complex128) complex128 {
	var z, prez complex128 = x / 3, 0

	for cmplx.Abs(z-prez) > 1e-10 {
		prez = z
		z = z - (z*z*z-x)/(z*z*3)
	}

	return z
}

func main() {
	fmt.Println(Cbrt(2))
	fmt.Println(cmplx.Pow(Cbrt(2), 3))
}
