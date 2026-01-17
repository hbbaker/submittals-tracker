package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var db *pgxpool.Pool

func main() {

	// Load env
	_ = godotenv.Load()

	dsn := os.Getenv("DB_URL")

	// Postgres Init
	ctx := context.Background()
	var err error
	db, err = pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close()
	fmt.Println("Connected successfully!")

	// Create Gin router
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Example endpoint that queries the database
	r.GET("/db-test", func(c *gin.Context) {
		var result string
		err := db.QueryRow(context.Background(), "SELECT 'Hello from PostgreSQL!'").Scan(&result)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": result})
	})

	r.Run(":3000")
}
