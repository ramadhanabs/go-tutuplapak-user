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
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database connection
	dbConn, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Repositories
	userRepo := repositories.NewUserRepository(dbConn)

	// Services
	authService := services.NewAuthService(userRepo, cfg)

	// Controllers
	authController := controllers.NewAuthController(authService)

	// Initialize router
	router := gin.Default()

	// Routes
	authRoutes := router.Group("/v1")
	{
		authRoutes.POST("/login/email", authController.LoginWithEmail)
		authRoutes.POST("/login/phone", authController.LoginWithPhone)
		authRoutes.POST("/register/email", authController.RegisterWithEmail)
		authRoutes.POST("/register/phone", authController.RegisterWithPhone)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)

	// Close the DB connection gracefully when the app exits
	defer db.CloseDB()
}
