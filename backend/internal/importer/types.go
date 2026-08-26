package importer

type IngredientData struct {
	Slug        string      `json:"slug"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Places      []PlaceData `json:"places"`
	Events      []EventData `json:"events"`
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

type EventData struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	TimePeriod  string   `json:"time_period"`
	Entity      string   `json:"entity"`
	Location    string   `json:"location"`
	Sources     []string `json:"sources"`
	Confidence  string   `json:"confidence"`
}