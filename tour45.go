package main

import (
	"fmt"
	"runtime"
)

// Switch
// http://go-tour-jp.appspot.com/#45

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}
