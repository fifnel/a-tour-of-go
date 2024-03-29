package main

import (
	"fmt"
	"math/cmplx"
)

// Basic types
// http://go-tour-jp.appspot.com/#14

var (
	ToBe   bool       = false
	MaxInt uint64     = 1 << 63
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
