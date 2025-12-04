// Package routes contains all the endpoints mapped to respective endpoints
package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jirugutema/gominimalapi/middlewares"
	"github.com/jirugutema/gominimalapi/controllers"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.Logger)

	//task related tasks
	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/ping", controllers.Ping)
	router.POST("/tasks", controllers.CreateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	//user related routes
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.POST("/users",controllers.CreateUser)
	router.DELETE("/users/:id",controllers.DeleteUser)
	router.POST("/auth/login", controllers.LoginHandler)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("error starting the server", err)
		return nil
	}
	return router
}
