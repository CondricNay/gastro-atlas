// package router

// import (
// 	"github.com/gin-gonic/gin"

// 	"github.com/CondricNay/gastro-atlas/internal/handlers"
// )

// func SetupRouter(ingredientHandler *handlers.IngredientHandler) *gin.Engine {
// 	r := gin.Default()

// 	r.GET("/ingredients", ingredientHandler.GetIngredients)
// 	r.GET("/ingredients/:slug", ingredientHandler.GetIngredientBySlug)

// 	return r
// }

package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/CondricNay/gastro-atlas/internal/handlers"
)

func SetupRouter(ingredientHandler *handlers.IngredientHandler) *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:5173",
	}

	r.Use(cors.New(config))

	r.GET("/ingredients", ingredientHandler.GetIngredients)
	r.GET("/ingredients/:slug", ingredientHandler.GetIngredientBySlug)

	return r
}