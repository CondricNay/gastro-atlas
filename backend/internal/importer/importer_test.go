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
		"events": [
			{
				"id": 1,
				"title": "Tomato reaches Europe",
				"description": "Tomatoes were introduced to Europe.",
				"time_period": "Late 15th century",
				"entity": "introduction",
				"location": "Spain",
				"sources": [],
				"confidence": "high"
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

	// Verify event
	events, err := queries.GetEventsByIngredient(ctx, ingredient.ID)
	if err != nil {
		t.Fatalf("get imported events: %v", err)
	}

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	event := events[0]

	if event.Title != "Tomato reaches Europe" {
		t.Errorf(
			"expected event title %q, got %q",
			"Tomato reaches Europe",
			event.Title,
		)
	}

	if event.Description != "Tomatoes were introduced to Europe." {
		t.Errorf(
			"expected event description %q, got %q",
			"Tomatoes were introduced to Europe.",
			event.Description,
		)
	}

	if event.TimePeriod != "Late 15th century" {
		t.Errorf(
			"expected time period %q, got %q",
			"Late 15th century",
			event.TimePeriod,
		)
	}

	if event.Entity != "introduction" {
		t.Errorf(
			"expected entity %q, got %q",
			"introduction",
			event.Entity,
		)
	}

	if event.Location != "Spain" {
		t.Errorf(
			"expected location %q, got %q",
			"Spain",
			event.Location,
		)
	}

	if len(event.Sources) != 0 {
		t.Errorf(
			"expected 0 sources, got %d",
			len(event.Sources),
		)
	}

	if event.Confidence != "high" {
		t.Errorf(
			"expected confidence %q, got %q",
			"high",
			event.Confidence,
		)
	}
}
