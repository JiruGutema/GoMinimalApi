//Package controllers handles http requrest for me
package controllers

import (

	"fmt"
	"strconv"
	"gominimalapi/data"
	"gominimalapi/models"
	"github.com/gin-gonic/gin"
)

func GetAllTasks(context *gin.Context) {
	context.JSON(200, data.Tasks)
}

func CreateTask(context *gin.Context) {
	var NewTask models.Task
	err := context.BindJSON(&NewTask)
	if err != nil {
		fmt.Println("error while binding the new task", err)
		return
	}

	// check if the task is already exist
	for i := 0; i < len(data.Tasks); i++ {
		if data.Tasks[i].Name == NewTask.Name {
			context.JSON(400, "Task already exists")
			return
		}
	}

	fmt.Println("New Task to be added:", NewTask)
	data.Tasks = append(data.Tasks, NewTask)
	context.JSON(201, data.Tasks)
}

func DeleteTask(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, "Invalid task ID")
	}


	idx := -1
	count := 0
	for count < len(data.Tasks) {
		if data.Tasks[count].ID == id {
			idx = count
			break
		}
		count++
	}
	if idx == -1 {
		context.JSON(404, "Task not found")
		return
	}

	data.Tasks = append(data.Tasks[:idx], data.Tasks[idx+1:]...)
	context.JSON(200, data.Tasks)
}

func GetHealthStatus(context *gin.Context) {
	context.JSON(200, "The server is up and running")
}
