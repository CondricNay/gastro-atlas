package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/CondricNay/gastro-atlas/internal/sqlc"
)

type IngredientHandler struct {
	queries *sqlc.Queries
}


func NewIngredientHandler(
	queries *sqlc.Queries,
) *IngredientHandler {

	return &IngredientHandler{
		queries: queries,
	}
}


func (h *IngredientHandler) GetIngredients(
	w http.ResponseWriter,
	r *http.Request,
) {

	ingredients, err :=
		h.queries.GetIngredients(
			r.Context(),
		)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			500,
		)
		return
	}


	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		ingredients,
	)
}