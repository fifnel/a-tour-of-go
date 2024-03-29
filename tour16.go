package main

import "fmt"

// Numeric Constants
// http://go-tour-jp.appspot.com/#16

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFload(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFload(Small))
	fmt.Println(needFload(Big))
}
