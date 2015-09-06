package main

import (
	"fmt"
	"image"
)

// Images
// http://go-tour-jp.appspot.com/#59

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
	fmt.Println(m.At(99, 99).RGBA())
	fmt.Println(m.At(99, 100).RGBA())
	fmt.Println(m.At(99, 101).RGBA())
}
