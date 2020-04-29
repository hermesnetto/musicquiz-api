package app

import (
	"log"
	"musicquiz-api/app/database"
	"musicquiz-api/controllers"
	"musicquiz-api/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// StartApplication starts a new application
func StartApplication() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	
	db, dbErr := database.Connect()
	if dbErr != nil {
		panic("Failed to connect database")
	}
	database.SetConnDatabase(db)
	defer db.Close()

	db.AutoMigrate(&models.User{})
	
	r := gin.Default()
	r.GET("/users/:id", controllers.HandleGetUser)
	r.POST("/users", controllers.HandleCreateUser)
	r.Run()
}
