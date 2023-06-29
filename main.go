package main

import (
	"net/http"
	"simpleGin/simplegin"
)

func main() {
	r := simplegin.New()
	r.GET("/", func(c *simplegin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello there!</h1>")
	})

	r.GET("/hello", func(c *simplegin.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *simplegin.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *simplegin.Context) {
		c.JSON(http.StatusOK, simplegin.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}

//func main() {
//	r := simplegin.New()
//	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
//		fmt.Fprintf(w, "hello there, it is version 2 of simplegin")
//	})
//	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
//		fmt.Fprintf(w, "this is index, it is version 2 of simplegin")
//	})
//	r.Run(":9999")
//}
