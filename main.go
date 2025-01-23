package main

import (
	"go-tutuplapak-user/config"
	"go-tutuplapak-user/controllers"
	"go-tutuplapak-user/db"
	"go-tutuplapak-user/repositories"
	"go-tutuplapak-user/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	dbConn, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	userRepo := repositories.NewUserRepository(dbConn)
	authService := services.NewAuthService(userRepo, cfg)
	authController := controllers.NewAuthController(authService)

	router := gin.Default()

	authRoutes := router.Group("/v1")
	{
		authRoutes.POST("/login/email", authController.LoginWithEmail)
		authRoutes.POST("/login/phone", authController.LoginWithPhone)
		authRoutes.POST("/register/email", authController.RegisterWithEmail)
		authRoutes.POST("/register/phone", authController.RegisterWithPhone)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)

	defer db.CloseDB()
}
