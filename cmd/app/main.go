package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/go_mysql/api/handler"
	"github.com/yourusername/go_mysql/api/router"
	"github.com/yourusername/go_mysql/config"
	"github.com/yourusername/go_mysql/internal/repository"
	"github.com/yourusername/go_mysql/internal/service"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	// Initialize database connection
	db, err := config.NewDatabase(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)

	// Setup router
	r := router.SetupRouter(userHandler)

	// Start server
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Starting server on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
