package router

import (
	"github.com/gin-gonic/gin"

	"github.com/CondricNay/gastro-atlas/internal/handlers"
)

func SetupRouter(ingredientHandler *handlers.IngredientHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/ingredients", ingredientHandler.GetIngredients)
	r.GET("/ingredients/:slug", ingredientHandler.GetIngredientBySlug)

	return r
}