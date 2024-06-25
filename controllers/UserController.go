package controllers

import (
	"crudApp/services"

	"github.com/gin-gonic/gin"
)

// Add User Controller
func AddUserController(c *gin.Context) {
	services.AddUserService(c)
}

// Remove User Controller
func RemoveUserController(c *gin.Context) {
	ch := make(chan bool)
	services.RemoveUserService(c, ch)
	<-ch
}

// Get User Controller
func GetUserById(c *gin.Context) {
	ch := make(chan bool)
	go services.GetUserByIdService(c, ch)
	<-ch
}

// Update User
func UpdateUser(c *gin.Context) {
	ch := make(chan bool)
	services.UpdateUserById(c, ch)
	<-ch
}
