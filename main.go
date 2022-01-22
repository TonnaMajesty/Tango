package main

import (
	tango "Tango/framework"
	"net/http"
)

func main() {
	r := tango.New()
	r.GET("/", func(c *tango.Context) {
		c.HTML(http.StatusOK, "<h1>Hello</h1>")
	})

	r.GET("/hello", func(c *tango.Context) {
		// expect /hello?name=xxx
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *tango.Context) {
		c.JSON(http.StatusOK, tango.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}