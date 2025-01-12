package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"sirapo/controllers"
	"sirapo/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	// CORS Configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:3000", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUsers)
	r.PATCH("/users/:id", controllers.UpdateUsers)
	r.DELETE("/users/:id", controllers.DeleteUsers)
	r.Run()
}
