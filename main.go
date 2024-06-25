package main

import (
	"fmt"

	"crudApp/controllers"

	"crudApp/config"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Go CURD APP Starting!")

	// Gin Router Define
	router := gin.Default()

	// REST Routes
	router.POST("/user/add", controllers.AddUserController)
	router.DELETE("/user/remove/:id", controllers.RemoveUserController)
	router.GET("/user/get/:id", controllers.GetUserById)
	router.PUT("/user/update/:id", controllers.UpdateUser)

	// Database Config Load
	config.DatabaseConfg()

	// Router Run Port
	router.SetTrustedProxies([]string{"172.30.1.158", "172.16.0.34", "127.0.0.1", "192.168.1.2", "10.0.0.0/8"})
	router.Run(":9001")
}
