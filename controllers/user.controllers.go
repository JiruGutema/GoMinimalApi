package controllers

import (
	"fmt"
	"strconv"

	"github.com/jirugutema/gominimalapi/data"
	"github.com/jirugutema/gominimalapi/dto"
	"github.com/jirugutema/gominimalapi/models"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(context *gin.Context) {
	users := data.Users
	if users == nil {
		context.JSON(200, []models.User{})
		return
	}
	context.JSON(200, users)
}

func CreateUser(context *gin.Context) {
	var user models.User
	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(400, "Invalid user data")
		return
	}

	length := len(data.Users)
	i := 0
	for i < length {
		if data.Users[i].Username == user.Username {
			context.JSON(400, "Username already exists")
		}
		i++
	}

	context.JSON(201, user)
}

func DeleteUser(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(400, "error deleting user")
	}

	exists := false
	length := len(data.Users)
	idx := 0

	for idx < length {
		if data.Users[idx].ID == id {
			exists = true
			break
		}
		idx++
	}

	if !exists {
		context.JSON(404, "user not found")
		return
	}

	data.Users = append(data.Users[:idx], data.Users[:idx+1]...)
	context.JSON(201, data.Users)
}

func GetUser(context *gin.Context) {
	var user models.User
	var id int

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(200, "Error getting user id")
		return
	}
	length := len(data.Users)
	i := 0
	for i < length {
		if data.Users[i].ID == id {
			user = data.Users[id]
		}
		i++
	}
	context.JSON(200, user)
}

func LoginHandler(context *gin.Context) {
	var login dto.Login

	err := context.BindJSON(&login)
	if err != nil {
		context.JSON(400, "error getting login information")
		return
	}

	fmt.Printf("Type: %T\n", login.ID)
	length := len(data.Users)
	idx := 0
	for idx < length {
		if data.Users[idx].Username == login.Username && data.Users[idx].ID == login.ID {
			context.JSON(200, gin.H{"logged": "success", "token": "fake_token"})
			return
		}
		idx++
	}
	context.JSON(404, "Invalid Username or ID")
}
