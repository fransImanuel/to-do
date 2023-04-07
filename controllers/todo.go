package controllers

import (
	"fmt"
	"net/http"
	"to-do/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Query("activity_group_id")
	var todos []models.Todo

	if id == "" {
		if err := db.Raw("select * from todos").Scan(&todos).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
			return
		}
	} else {
		if err := db.Raw("select * from todos where activity_group_id = ?", id).Scan(&todos).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todos})
}

func GetTodoByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	param := c.Param("id")
	var todos models.Todo

	if err := db.Raw("select * from todos where todo_id = ?", param).Scan(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	if (todos == models.Todo{}) {
		c.JSON(http.StatusOK, gin.H{"status": "Not Found", "message": "Todo with ID 8 Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todos})
}

func PostTodo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var todos models.Todo
	var B models.PostTodo_Req

	if err := c.ShouldBindJSON(&B); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": "title cannot be null"})
		return
	}

	if err := db.Raw("insert into todos(title, activity_group_id, is_active, priority) values(?,?,?,?)", B.Title, B.Activity_Group_Id, B.Is_Active, B.Priority).Scan(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	db.Last(&todos)

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": todos})
}

func UpdateTodoByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var todo models.Todo
	var B models.PostTodo_Req

	if err := c.ShouldBindJSON(&B); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request", "message": "title cannot be null"})
		return
	}

	if err := db.Exec("update todos set title = ?, activity_group_id = ?, priority = ?, is_active = ? where todo_id = ?", B.Title, B.Activity_Group_Id, B.Is_Active, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": http.StatusInternalServerError})
		return
	}

	db.First(&todo, "todo_id = ?", id)

	if (todo == models.Todo{}) {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Todo with ID %s Not Found", id), "status": "Not Found"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Success", "message": "Success", "data": todo})
}

func DeleteTodoByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	result := db.Exec("delete from todo where todo_id = ?", id)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": http.StatusInternalServerError})
		return
	}

	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Todo with ID %s Not Found", id), "status": "Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Success", "message": "Success", "data": struct{}{}})
}
