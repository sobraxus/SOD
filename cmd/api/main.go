package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/sobraxus/SOD/internal/db"
	"github.com/sobraxus/SOD/internal/handlers"
	"github.com/sobraxus/SOD/internal/repositories"
)

func main() {
	ctx := context.Background()

	// Initialize PostgreSQL
	connStr := "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable"
	database, err := db.NewPostgres(ctx, connStr)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer database.Pool.Close()

	// Create repository and handler
	pgxConn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Failed to get PGX connection: %v", err)
	}
	caseRepo := repositories.NewCaseRepository(pgxConn)
	caseHandler := handlers.NewCaseHandler(caseRepo)

	// Configure Gin
	r := gin.Default()

	// Routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Case endpoints
	caseGroup := r.Group("/cases")
	{
		caseGroup.POST("", caseHandler.CreateCase)
		caseGroup.GET("/:id", caseHandler.GetCaseByID)
	}

	// Start server
	log.Println("Server running on :8080")
	r.Run(":8080")
}
