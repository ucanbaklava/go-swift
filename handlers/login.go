package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ucanbaklava/go-auth/database"
	"github.com/ucanbaklava/go-auth/models"
	"github.com/ucanbaklava/go-auth/utils"
)

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by username or email
	var user models.User

	query := `SELECT id, username, email, password, created_at, updated_at, role, is_active FROM users WHERE username = ?`
	err := database.DB.QueryRow(query, input.Username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.Role, &user.IsActive)
	fmt.Print(err)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user"})
		}
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := utils.GenerateJWT(user.Username, user.Role, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	user.LastLogin = time.Now()
	updateQuery := `UPDATE users SET last_login = ? WHERE id = ?`
	_, err = database.DB.Exec(updateQuery, user.LastLogin, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update last login"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
