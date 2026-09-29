package importer

type IngredientData struct {
	Slug        string      `json:"slug"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Events      []EventData `json:"events"`
}

type EventData struct {
	Description string   `json:"description"`
	TimePeriod  string   `json:"time_period"`
	Location    string   `json:"location"`
	StartYear   *int32   `json:"start_year"`
	EndYear     *int32   `json:"end_year"`
	Latitude    *float64 `json:"latitude"`
	Longitude   *float64 `json:"longitude"`
}