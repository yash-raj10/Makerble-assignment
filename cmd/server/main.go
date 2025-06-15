package main

import (
	"log"
	"os"

	"ass4/internal/config"
	"ass4/internal/handler"
	"ass4/internal/models"
	"ass4/internal/routes"
	"ass4/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load() 
	cfg := config.LoadConfig()
	utils.SetJWTKey(cfg.JWTSecret)
	db, err := gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&models.User{}, &models.Patient{})

	authHandler := handler.NewAuthHandler(db)
	patientHandler := handler.NewPatientHandler(db)

	r := gin.Default()
	routes.RegisterRoutes(r, authHandler, patientHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}
