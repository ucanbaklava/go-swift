package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/ucanbaklava/go-auth/database"
	"github.com/ucanbaklava/go-auth/handlers"
	"github.com/ucanbaklava/go-auth/middleware"
	"github.com/ucanbaklava/go-auth/repository"
)

func main() {
	database.InitDatabase()

	db, _ := sqlx.Connect("sqlite3", "./test.db")

	postRepo := repository.NewPostRepository(db)

	r := gin.Default()

	r.POST("/api/register", handlers.Register)
	r.POST("/api/login", handlers.Login)

	r.GET("/api/test", func(c *gin.Context) {
		post, err := postRepo.GetByID(1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, post)
	})

	auth := r.Group("/api/auth")

	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/protected", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "you are authorized"})
		})
		auth.GET("/admin", middleware.AdminMiddleware(), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "you are an admin"})
		})
	}

	r.Run(":8080")

}
