package importer

type IngredientData struct {
	Slug        string      `json:"slug"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Events      []EventData `json:"events"`
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