package main

// Exercise: Slices
// http://go-tour-jp.appspot.com/#35

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for y := range image {
		image[y] = make([]uint8, dx)
	}

	for y := range image {
		for x := range image[y] {
			image[y][x] = uint8((x + y) / 2)
		}
	}

	return image
}

func main() {
	pic.Show(Pic)
}
