package main

import (
   "fmt"
   "net/http"
   "github.com/gin-gonic/gin"
)

var version string = "0.0.0"

func main() {
	r := gin.Default()
        r.GET("/info", func(c *gin.Context) {
          c.JSON(http.StatusOK, gin.H{
            "name": "gocode",
	    "author": "liman",
	    "version": version,
          })
        })
	fmt.Println("Server running on version", version)
	r.Run()
}
