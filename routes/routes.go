// Package routes contains all the endpoints mapped to respective endpoints
package routes

import (
	"fmt"

	"gominimalapi/controllers"
	"gominimalapi/middlewares"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.Logger)

	//task related tasks
	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/health", controllers.GetHealthStatus)
	router.POST("/tasks", controllers.CreateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	//user related routes
	router.GET("users", controllers.GetAllUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.POST("/users",controllers.CreateUser)
	router.DELETE("/users/:id",controllers.DeleteUser)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("error starting the server", err)
		return nil
	}
	return router
}
