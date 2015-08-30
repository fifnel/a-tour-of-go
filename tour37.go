package main

import "fmt"

// Maps
// http://go-tour-jp.appspot.com/#37

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var mm = map[string]string{
		"a": "A",
		"b": "B",
	}
	fmt.Println(mm)
}
