// import "github.com/gin-gonic/gin"

// func main() {
//   gin.SetMode(gin.ReleaseMode) //optional to not get warning
//   router := gin.Default()
//   router.GET("/ping", func(c *gin.Context) {
//     c.JSON(200, gin.H{
//       "message": "pong",
//     })
//   })
//   router.Run() // listens on 0.0.0.0:8080 by default
// }

package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/CondricNay/gastro-atlas/internal/database"
	"github.com/CondricNay/gastro-atlas/internal/handlers"
	"github.com/CondricNay/gastro-atlas/internal/sqlc"

// 	"github.com/CondricNay/gastro-atlas/internal/models"

	"github.com/joho/godotenv"
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


	ingredientHandler :=
		handlers.NewIngredientHandler(queries)

	r := chi.NewRouter()

	r.Get("/ingredients", ingredientHandler.GetIngredients)

	log.Println("Server running on :8080")

	http.ListenAndServe(":8080", r)
}