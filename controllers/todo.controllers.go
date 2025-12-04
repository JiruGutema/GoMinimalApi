// Package controllers handles http requrest for me
package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jirugutema/gominimalapi/config"
	"github.com/jirugutema/gominimalapi/data"
	"github.com/jirugutema/gominimalapi/models"
)

func GetAllTasks(context *gin.Context) {
	query := "SELECT id, name, isCompleted FROM todo"
	rows, err := config.DB.Query(query)
	if err != nil {
		context.JSON(400, gin.H{
			"success": false,
			"message": "error in reading the database",
		})
	}

	defer rows.Close()
	var results []models.Task
	for rows.Next() {
		var id int
		var name string
		var isCompleted bool

		err := rows.Scan(&id, &name, &isCompleted)
		if err != nil {

			context.JSON(400, gin.H{
				"success": false,
				"message": "error in reading the database",
			})
			return

		}
		newTask := models.Task{
			ID:          id,
			Name:        name,
			IsCompleted: isCompleted,
		}
		results = append(results, newTask)
	}

	context.JSON(200, results)
}

func CreateTask(context *gin.Context) {
	var NewTask models.Task
	err := context.BindJSON(&NewTask)
	if err != nil {
		fmt.Println("error while binding the new task", err)
		return
	}

	query := "INSERT INTO todo (ID, name, isCompleted) VALUES ($1, $2, $3) RETURNING id"
	row := config.DB.QueryRow(query, NewTask.ID, NewTask.Name, NewTask.IsCompleted)
	e := row.Scan(&NewTask.ID)
	if e != nil {
		context.JSON(200, gin.H{
			"success": false,
			"message": "can't create a specified task",
		})
		return
	}

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
