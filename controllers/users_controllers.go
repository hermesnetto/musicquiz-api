package controllers

import (
	"musicquiz-api/models"
	"musicquiz-api/services"
	"musicquiz-api/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandleGetUser gets a user
func HandleGetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("User ID should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, err := services.GetUser(userID)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusCreated, user)
}

// HandleCreateUser creates a new user
func HandleCreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	createdUser, err := services.CreateUser(&user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}
