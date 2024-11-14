package main

import (
	"webclinicalrecords/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./templates/*")

	r.GET("/ping", internal.PongRoute)

	r.GET("/", internal.HomeRoute)

	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}
