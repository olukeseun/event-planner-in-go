package main

import (
	"event-planner-in-go/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type registerRquest struct {
	FirstName   string `json:"firstname" binding:"required"`
	LastName    string `json:"lastname" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Passowrd    string `json:"-"`
	PhoneNumber string `json:"phonenumber"`
}

func (app *application) register(c *gin.Context) {
	var register registerRquest

	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(register.Passowrd), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	register.Passowrd = string(hashPassword)

	user := database.User{
		FirstName:     register.FirstName,
		LastName:      register.LastName,
		Email:         register.Email,
		Password_hash: register.Passowrd,
		PhoneNumber:   register.PhoneNumber,
	}

	err = app.models.Users.Insert(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
