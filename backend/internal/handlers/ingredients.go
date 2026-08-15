package handlers

import (
	"github.com/CondricNay/gastro-atlas/internal/sqlc"
	"github.com/gin-gonic/gin"
)

type IngredientHandler struct {
	queries *sqlc.Queries
}

func NewIngredientHandler(queries *sqlc.Queries) *IngredientHandler {
	return &IngredientHandler{queries: queries}
}

func (h *IngredientHandler) GetIngredients(c *gin.Context) {
	ingredients, err :=
		h.queries.GetIngredients(
			c.Request.Context(),
		)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, ingredients)
}

func (h *IngredientHandler) GetIngredientBySlug(c *gin.Context) {
	slug := c.Param("slug")

	ingredient, err := h.queries.GetIngredientBySlug(
		c.Request.Context(), slug,
	)

	if err != nil {
		c.JSON(404, gin.H{"error": "ingredient not found"})
		return
	}

	places, err := h.queries.GetIngredientPlaces(
		c.Request.Context(),
		ingredient.ID,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"id":          ingredient.ID,
		"slug":        ingredient.Slug,
		"name":        ingredient.Name,
		"description": ingredient.Description,
		"places":      places,
	})
}
