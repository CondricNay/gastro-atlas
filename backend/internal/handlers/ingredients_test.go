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

	placeID, err := queries.UpsertPlace(
		ctx,
		sqlc.UpsertPlaceParams{
			Name:      "Test Mexico",
			Type:      "origin",
			Latitude:  19.4326,
			Longitude: -99.1332,
		},
	)
	if err != nil {
		t.Fatalf("insert test place: %v", err)
	}

	err = queries.UpsertIngredientPlace(
		ctx,
		sqlc.UpsertIngredientPlaceParams{
			IngredientID: ingredientID,
			PlaceID:      placeID,
			Relationship: "origin",
			StartYear: pgtype.Int4{
				Int32: 1500,
				Valid: true,
			},
			EndYear: pgtype.Int4{
				Int32: 1600,
				Valid: true,
			},
			Notes: pgtype.Text{
				String: "Test place",
				Valid:  true,
			},
		},
	)
	if err != nil {
		t.Fatalf("insert ingredient place: %v", err)
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
		Places []struct {
			Name string `json:"name"`
		} `json:"places"`
	}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf(
			"response is not valid JSON: %v",
			err,
		)
	}

	if response.Slug != "test-tomato" {
		t.Errorf(
			"expected slug %q, got %q",
			"test-tomato",
			response.Slug,
		)
	}

	if len(response.Places) != 1 {
		t.Fatalf(
			"expected 1 place, got %d",
			len(response.Places),
		)
	}

	if response.Places[0].Name != "Test Mexico" {
		t.Errorf(
			"expected place %q, got %q",
			"Test Mexico",
			response.Places[0].Name,
		)
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