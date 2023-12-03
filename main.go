package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.Static("/assets", "./assets")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", func() {})
	})

	router.Run()

}
