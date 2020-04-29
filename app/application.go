package app

import (
	"musicquiz-api/controllers"

	"github.com/gin-gonic/gin"
)

// StartApplication starts a new application
func StartApplication() {
	r := gin.Default()
	r.GET("/users/:id", controllers.HandleGetUser)
	r.POST("/users", controllers.HandleCreateUser)
	r.Run()
}
