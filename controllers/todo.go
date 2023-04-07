package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Query("activity_group_id")

	if id == "" {
		db.Raw("select * from todos")
	} else {

	}
}
