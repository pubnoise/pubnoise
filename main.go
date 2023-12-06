package main

import (
	"fmt"
	"pubnoise/db"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		err = fmt.Errorf("%s", err)
		fmt.Println(err)
	}
	db.Connection()

	// router := gin.Default()
	// router.LoadHTMLGlob("views/*")
	// router.Static("/assets", "./assets")

	// router.GET("/", func(c *gin.Context) {

	// 	c.HTML(http.StatusOK, "index.html", func() {})
	// })

	// router.Run()

}
