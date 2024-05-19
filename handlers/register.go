package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ucanbaklava/go-auth/constants"
	"github.com/ucanbaklava/go-auth/database"
	"github.com/ucanbaklava/go-auth/models"
	"github.com/ucanbaklava/go-auth/utils"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		//TODO: dont send this error to the client
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	fmt.Println("password", user.Password)
	fmt.Println("hash error", err)
	fmt.Println("hashed", hashedPassword)

	user.Password = hashedPassword
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Role = string(constants.User)

	query := `INSERT INTO users (username, email, password, created_at, updated_at, role, is_active) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err = database.DB.Exec(query, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.Role, user.IsActive)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}
