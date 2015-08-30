package main

import "fmt"

// Making slices
// http://go-tour-jp.appspot.com/#32

func main() {
	// 配列は値
	// array := [3]int{1, 2, 3}

	// スライスは参照
	// slice := []int{1,2,3}
	// make(型, 長さ, キャパシティ)

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	bb := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	c := bb[:2] // 0番目〜2番目までの要素
	printSlice("c", c)

	d := bb[2:5] // 2番目〜5-1番目の要素
	printSlice("d", d)

	e := bb[3:] // 3番目〜最後の要素
	printSlice("e", e)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
