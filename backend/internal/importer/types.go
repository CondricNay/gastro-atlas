package importer

type IngredientData struct {
	Slug        string      `json:"slug"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Places      []PlaceData `json:"places"`
}

type PlaceData struct {
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Relationship string  `json:"relationship"`
	StartYear    *int    `json:"startYear"`
	EndYear      *int    `json:"endYear"`
	Notes        string  `json:"notes"`
}
