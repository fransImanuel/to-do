package controllers

import "github.com/gin-gonic/gin"

func HelloWorldController(c *gin.Context) {
	c.JSON(200, "Hello World")
}
