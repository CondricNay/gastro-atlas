package importer

import "fmt"

func ValidateEvent(event EventData) error {
	if event.Title == "" {
		return fmt.Errorf("title is required")
	}

	return nil
}

func ValidateIngredient(ingredient IngredientData) error {
	for _, event := range ingredient.Events {
		if err := ValidateEvent(event); err != nil {
			return fmt.Errorf("invalid event %q: %w", event.Title, err)
		}
	}

	return nil
}
