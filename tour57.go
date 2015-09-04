package main

import (
	"fmt"
	"net/http"
)

// Web servers
// http://go-tour-jp.appspot.com/#57

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}
