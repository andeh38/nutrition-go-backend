package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"nutrition-api/controller"
	"nutrition-api/database"
	"nutrition-api/model"
	"nutrition-api/middleware"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Day{})
	database.Database.AutoMigrate(&model.Food{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.CORSMiddleware())
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/food", controller.AddFood)
	protectedRoutes.GET("/food/", controller.FindFoodByName)
	protectedRoutes.GET("/food/all", controller.GetAllFood)

	protectedRoutes.POST("/days/search", controller.GetDay)
	protectedRoutes.GET("/days/all", controller.AllDays)
	protectedRoutes.PUT("/days/update", controller.Update)

	protectedRoutes.GET("/analytics", controller.GetAnalytics)
	
	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
