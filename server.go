package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLTemplates( "test.html" )

	r.GET( "/", func( c *gin.Context ) {
		c.HTML( 200, "test.html", gin.H{} )
	})

	r.Run(":8080")
}
