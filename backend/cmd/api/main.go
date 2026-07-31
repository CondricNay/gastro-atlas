package main

import (
	"log"

	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/CondricNay/gastro-atlas/internal/database"
	"github.com/CondricNay/gastro-atlas/internal/handlers"
	"github.com/CondricNay/gastro-atlas/internal/router"
	"github.com/CondricNay/gastro-atlas/internal/sqlc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	pool, err := database.New()

	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	queries := sqlc.New(pool)

	ingredientHandler := handlers.NewIngredientHandler(queries)
    r := router.SetupRouter(ingredientHandler)

	log.Println("Server running on :8080")

	r.Run(":8080")
}