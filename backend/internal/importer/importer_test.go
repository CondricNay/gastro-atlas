package importer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/CondricNay/gastro-atlas/internal/sqlc"
	"github.com/CondricNay/gastro-atlas/internal/testutil"
)

func TestImportIngredient(t *testing.T) {
	db, ctx := testutil.SetupTestDB(t)

	dir := t.TempDir()
	filePath := filepath.Join(dir, "test-ingredient.json")

	jsonData := `{
		"name": "Test Tomato",
		"slug": "test-tomato",
		"description": "A test ingredient",
		"places": [
			{
				"name": "Test Mexico",
				"type": "origin",
				"latitude": 19.4326,
				"longitude": -99.1332,
				"relationship": "origin",
				"startYear": 1500,
				"endYear": 1600,
				"notes": "Test place"
			}
		]
	}`

	if err := os.WriteFile(filePath, []byte(jsonData), 0644); err != nil {
		t.Fatalf("write test JSON: %v", err)
	}

	if err := importIngredient(ctx, db, filePath); err != nil {
		t.Fatalf("import ingredient: %v", err)
	}

	queries := sqlc.New(db)

	// Verify ingredient
	ingredient, err := queries.GetIngredientBySlug(ctx, "test-tomato")
	if err != nil {
		t.Fatalf("get imported ingredient: %v", err)
	}

	if ingredient.Name != "Test Tomato" {
		t.Errorf("expected name %q, got %q", "Test Tomato", ingredient.Name)
	}

	if ingredient.Slug != "test-tomato" {
		t.Errorf("expected slug %q, got %q", "test-tomato", ingredient.Slug)
	}

	if !ingredient.Description.Valid {
		t.Fatal("expected description to be present")
	}

	if ingredient.Description.String != "A test ingredient" {
		t.Errorf(
			"expected description %q, got %q",
			"A test ingredient",
			ingredient.Description.String,
		)
	}

	// Verify places and ingredient_places
	places, err := queries.GetIngredientPlaces(ctx, ingredient.ID)
	if err != nil {
		t.Fatalf("get imported ingredient places: %v", err)
	}

	if len(places) != 1 {
		t.Fatalf("expected 1 place, got %d", len(places))
	}

	place := places[0]

	if place.Name != "Test Mexico" {
		t.Errorf("expected place name %q, got %q", "Test Mexico", place.Name)
	}

	if place.Type != "origin" {
		t.Errorf("expected place type %q, got %q", "origin", place.Type)
	}

	if place.Latitude != 19.4326 {
		t.Errorf("expected latitude %f, got %f", 19.4326, place.Latitude)
	}

	if place.Longitude != -99.1332 {
		t.Errorf("expected longitude %f, got %f", -99.1332, place.Longitude)
	}

	if place.Relationship != "origin" {
		t.Errorf(
			"expected relationship %q, got %q",
			"origin",
			place.Relationship,
		)
	}

	if !place.StartYear.Valid || place.StartYear.Int32 != 1500 {
		t.Errorf("expected start year 1500, got %+v", place.StartYear)
	}

	if !place.EndYear.Valid || place.EndYear.Int32 != 1600 {
		t.Errorf("expected end year 1600, got %+v", place.EndYear)
	}

	if !place.Notes.Valid || place.Notes.String != "Test place" {
		t.Errorf("expected notes %q, got %+v", "Test place", place.Notes)
	}
}

func TestImportIngredient_Idempotent(t *testing.T) {
	db, ctx := testutil.SetupTestDB(t)

	dir := t.TempDir()
	filePath := filepath.Join(dir, "test-ingredient.json")

	jsonData := `{
		"name": "Test Tomato",
		"slug": "test-tomato",
		"description": "A test ingredient",
		"places": [
			{
				"name": "Test Mexico",
				"type": "origin",
				"latitude": 19.4326,
				"longitude": -99.1332,
				"relationship": "origin",
				"startYear": 1500,
				"endYear": 1600,
				"notes": "Test place"
			}
		]
	}`

	if err := os.WriteFile(filePath, []byte(jsonData), 0644); err != nil {
		t.Fatalf("write test JSON: %v", err)
	}

	// Import the same ingredient twice.
	if err := importIngredient(ctx, db, filePath); err != nil {
		t.Fatalf("first import: %v", err)
	}

	if err := importIngredient(ctx, db, filePath); err != nil {
		t.Fatalf("second import: %v", err)
	}

	queries := sqlc.New(db)

	// Verify only one ingredient exists.
	var ingredientCount int

	err := db.QueryRow(
		ctx,
		`SELECT COUNT(*)
		 FROM ingredients
		 WHERE slug = $1`,
		"test-tomato",
	).Scan(&ingredientCount)

	if err != nil {
		t.Fatalf("count ingredients: %v", err)
	}

	if ingredientCount != 1 {
		t.Errorf(
			"expected 1 ingredient, got %d",
			ingredientCount,
		)
	}

	// Verify only one place relationship exists.
	ingredient, err := queries.GetIngredientBySlug(ctx, "test-tomato")
	if err != nil {
		t.Fatalf("get ingredient: %v", err)
	}

	places, err := queries.GetIngredientPlaces(ctx, ingredient.ID)
	if err != nil {
		t.Fatalf("get ingredient places: %v", err)
	}

	if len(places) != 1 {
		t.Errorf(
			"expected 1 ingredient-place relationship, got %d",
			len(places),
		)
	}
}