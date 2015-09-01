package main

import "fmt"

func fibonacci() func() int {
	n0 := 0
	n1 := 1
	return func() int {
		n2 := n0 + n1
		n0 = n1
		n1 = n2
		return n2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
