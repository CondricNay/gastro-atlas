package handlers

import (
	"encoding/json"
	"net/http"
)

type Ingredient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetIngredients(w http.ResponseWriter, r *http.Request) {
	ingredients := []Ingredient{
		{ID: 1, Name: "Tomato"},
		{ID: 2, Name: "Potato"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ingredients)
}