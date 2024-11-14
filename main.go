package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "helloworld the path is%q\n", c.Path)
	})
	r.GET("/hello", func(c *gee.Context) {
		h := gee.H{}
		for k, v := range c.Req.Header {
			//fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			h[k] = v
		}
		c.JSON(http.StatusOK, h)
	})
	r.POST("/login", func(c *gee.Context) {
		userid := c.Query("userid")
		password := c.Query("password")
		c.JSON(http.StatusOK, gee.H{
			"userid":   userid,
			"password": password,
		})
	})
	r.Run(":9999")
}
