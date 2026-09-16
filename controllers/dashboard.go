package controllers

import "github.com/gin-gonic/gin"

func Dashboard(c *gin.Context) {

	c.HTML(200, "dashboard.html", nil)
}
