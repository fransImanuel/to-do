package server

import (
	"to-do/controllers"
	"to-do/db"

	"github.com/gin-gonic/gin"
)

func NewRouter(mysql *db.MysqlConn) *gin.Engine {
	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Set("db", mysql.GetDBInstance()) //set mysql db instance
	})

	health := new(controllers.HealthController)

	r.GET("/health", health.Status)
	// r.Use(middlewares.AuthMiddleware())

	v1 := r.Group("v1")
	{
		userGroup := v1.Group("hello")
		{
			userGroup.GET("/:id", controllers.HelloWorldController)
		}
	}

	//Activity
	r.GET("/activity-groups", controllers.GetAllActivity)
	r.GET("/activity-groups/:id", controllers.GetActivityByID)
	r.POST("/activity-groups", controllers.PostActivity)
	r.PATCH("/activity-groups/:id", controllers.UpdateActivityByID)
	r.DELETE("/activity-groups/:id", controllers.DeleteActivityByID)

	//Todo
	r.GET("/todo-items", controllers.GetAllTodo)

	return r

}
