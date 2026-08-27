package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/CondricNay/gastro-atlas/internal/handlers"
	"github.com/CondricNay/gastro-atlas/internal/router"
	"github.com/CondricNay/gastro-atlas/internal/sqlc"
	"github.com/CondricNay/gastro-atlas/internal/testutil"
)

func setupTestRouter(
	t *testing.T,
) (*gin.Engine, context.Context, *sqlc.Queries) {
	t.Helper()

	db, ctx := testutil.SetupTestDB(t)
	queries := sqlc.New(db)

	handler := handlers.NewIngredientHandler(queries)
	r := router.SetupRouter(handler)

	return r, ctx, queries
}

func seedTestIngredient(
	t *testing.T,
	ctx context.Context,
	queries *sqlc.Queries,
) {
	t.Helper()

	ingredientID, err := queries.UpsertIngredient(
		ctx,
		sqlc.UpsertIngredientParams{
			Name: "Test Tomato",
			Slug: "test-tomato",
			Description: pgtype.Text{
				String: "A test ingredient",
				Valid:  true,
			},
		},
	)
	if err != nil {
		t.Fatalf("insert test ingredient: %v", err)
	}

	_, err = queries.CreateEvent(
		ctx,
		sqlc.CreateEventParams{
			IngredientID: ingredientID,
			Title:        "Tomato reaches Europe",
			Description:  "Tomatoes were introduced to Europe.",
			TimePeriod:   "Late 15th century",
			Entity:       "introduction",
			Location:     "Spain",
			Sources:      []string{},
			Confidence:   "high",
		},
	)
	if err != nil {
		t.Fatalf("insert test event: %v", err)
	}
}

func TestGetIngredients(t *testing.T) {
	r, ctx, queries := setupTestRouter(t)

	seedTestIngredient(t, ctx, queries)

	req := httptest.NewRequest(
		http.MethodGet,
		"/ingredients",
		nil,
	)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			rec.Code,
		)
	}

	var response []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf(
			"response is not valid JSON: %v",
			err,
		)
	}

	if len(response) != 1 {
		t.Fatalf(
			"expected 1 ingredient, got %d",
			len(response),
		)
	}

	if response[0].Name != "Test Tomato" {
		t.Errorf(
			"expected name %q, got %q",
			"Test Tomato",
			response[0].Name,
		)
	}

	if response[0].Slug != "test-tomato" {
		t.Errorf(
			"expected slug %q, got %q",
			"test-tomato",
			response[0].Slug,
		)
	}
}

func TestGetIngredientBySlug(t *testing.T) {
	r, ctx, queries := setupTestRouter(t)

	seedTestIngredient(t, ctx, queries)

	req := httptest.NewRequest(
		http.MethodGet,
		"/ingredients/test-tomato",
		nil,
	)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			rec.Code,
		)
	}

	var response struct {
		Name   string `json:"name"`
		Slug   string `json:"slug"`
		Events []struct {
			ID          int      `json:"id"`
			Title       string   `json:"title"`
			Description string   `json:"description"`
			TimePeriod  string   `json:"timePeriod"`
			Entity      string   `json:"entity"`
			Location    string   `json:"location"`
			Sources     []string `json:"sources"`
			Confidence  string   `json:"confidence"`
		} `json:"events"`
	}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf(
			"response is not valid JSON: %v",
			err,
		)
	}

	if response.Name != "Test Tomato" {
		t.Errorf(
			"expected name %q, got %q",
			"Test Tomato",
			response.Name,
		)
	}

	if response.Slug != "test-tomato" {
		t.Errorf(
			"expected slug %q, got %q",
			"test-tomato",
			response.Slug,
		)
	}

	if len(response.Events) != 1 {
		t.Fatalf(
			"expected 1 event, got %d",
			len(response.Events),
		)
	}

	event := response.Events[0]

	if event.ID == 0 {
		t.Error("expected event ID to be non-zero")
	}

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

	if event.Confidence != "high" {
		t.Errorf(
			"expected confidence %q, got %q",
			"high",
			event.Confidence,
		)
	}

	if event.Sources == nil {
		t.Error("expected sources to be non-nil")
	}
}

func TestGetIngredientBySlugNotFound(t *testing.T) {
	r, _, _ := setupTestRouter(t)

	req := httptest.NewRequest(
		http.MethodGet,
		"/ingredients/does-not-exist",
		nil,
	)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status 404, got %d",
			rec.Code,
		)
	}
}