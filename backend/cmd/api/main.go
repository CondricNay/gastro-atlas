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

// package main

// import (
// 	"context"
// 	"log"

// 	"github.com/CondricNay/gastro-atlas/internal/database"
// 	"github.com/CondricNay/gastro-atlas/internal/models"

// 	"github.com/joho/godotenv"
// )

// func main() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Failed to load .env")
// 	}

// 	db, err := database.New()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	log.Println("✅ Connected to PostgreSQL!")

// 	rows, err := db.Query(
// 		context.Background(),
// 		`
// 		SELECT id, name, slug, description
// 		FROM ingredients
// 		ORDER BY id
// 		`,
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var ingredient models.Ingredient

// 		err := rows.Scan(
// 			&ingredient.ID,
// 			&ingredient.Name,
// 			&ingredient.Slug,
// 			&ingredient.Description,
// 		)

// 		if err != nil {
// 			log.Fatal(err)
// 		}

//		log.Println(ingredient)
//	}
//
// }
package main

import (
	"log"
	"net/http"

	"github.com/CondricNay/gastro-atlas/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	r.Get("/ingredients", handlers.GetIngredients)

	log.Println("Server running on :8080")

	http.ListenAndServe(":8080", r)
}