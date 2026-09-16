package main

import (
	"ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("view/*")

	r.GET("/", controllers.Dashboard)

	r.Run(":8080")

}
