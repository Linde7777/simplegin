package main

import (
	"SimpleGin/simplegin"
	"fmt"
	"net/http"
)

func main() {
	r := simplegin.New()
	r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello there, it is version 2 of simplegin")
	})
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "this is index, it is version 2 of simplegin")
	})
	r.Run(":9999")
}
