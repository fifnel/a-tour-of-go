package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)

	counter := make(map[string]int)
	for _, w := range words {
		counter[w]++
	}
	return counter
}

func main() {
	wc.Test(WordCount)
}
