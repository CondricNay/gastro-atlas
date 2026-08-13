package importer

import "fmt"

func ValidatePlace(place PlaceData) error {
	if place.Latitude < -90 || place.Latitude > 90 {
		return fmt.Errorf("latitude must be between -90 and 90")
	}

	if place.Longitude < -180 || place.Longitude > 180 {
		return fmt.Errorf("longitude must be between -180 and 180")
	}

	// Skip nil checks for now
	if place.StartYear != nil && place.EndYear != nil && 
		*place.StartYear > *place.EndYear {
		return fmt.Errorf("start year must be before or equal to end year")
	}

	return nil
}

func ValidateIngredient(ingredient IngredientData) error {
	for _, place := range ingredient.Places {
		if err := ValidatePlace(place); err != nil {
			return fmt.Errorf("invalid place %q: %w", place.Name, err)
		}
	}

	return nil
}