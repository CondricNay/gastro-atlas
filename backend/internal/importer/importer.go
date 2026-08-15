package importer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/CondricNay/gastro-atlas/internal/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(ctx context.Context, db *pgxpool.Pool, dataDir string) error {
	files, err := filepath.Glob(filepath.Join(dataDir, "*.json"))
	if err != nil {
		return fmt.Errorf("find ingredient files: %w", err)
	}

	for _, file := range files {
		if err := importIngredient(ctx, db, file); err != nil {
			return fmt.Errorf("import %s: %w", file, err)
		}
	}

	fmt.Printf("Imported %d ingredient(s)\n", len(files))

	return nil
}

func importIngredient(
	ctx context.Context,
	db *pgxpool.Pool,
	filePath string,
) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read JSON: %w", err)
	}

	var ingredient IngredientData

	if err := json.Unmarshal(data, &ingredient); err != nil {
		return fmt.Errorf("parse JSON: %w", err)
	}

	if err := ValidateIngredient(ingredient); err != nil {
		return fmt.Errorf("validate data: %w", err)
	}

	fmt.Printf("Importing %s...\n", ingredient.Name)

	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer tx.Rollback(ctx)

	queries := sqlc.New(db).WithTx(tx)

	ingredientID, err := queries.UpsertIngredient(
		ctx,
		sqlc.UpsertIngredientParams{
			Name: ingredient.Name,
			Slug: ingredient.Slug,
			Description: pgtype.Text{
				String: ingredient.Description,
				Valid:  ingredient.Description != "",
			},
		},
	)
	if err != nil {
		return fmt.Errorf("upsert ingredient: %w", err)
	}

	for _, place := range ingredient.Places {
		placeID, err := queries.UpsertPlace(
			ctx,
			sqlc.UpsertPlaceParams{
				Name:      place.Name,
				Type:      place.Type,
				Latitude:  place.Latitude,
				Longitude: place.Longitude,
			},
		)
		if err != nil {
			return fmt.Errorf("upsert place %q: %w", place.Name, err)
		}

		if err := queries.UpsertIngredientPlace(
			ctx,
			sqlc.UpsertIngredientPlaceParams{
				IngredientID: ingredientID,
				PlaceID:      placeID,
				Relationship: place.Relationship,
				StartYear:    intToPgtype(place.StartYear),
				EndYear:      intToPgtype(place.EndYear),
				Notes: pgtype.Text{
					String: place.Notes,
					Valid:  place.Notes != "",
				},
			},
		); err != nil {
			return fmt.Errorf(
				"upsert relationship for %q: %w",
				place.Name,
				err,
			)
		}

		fmt.Printf("✓ %s\n", place.Name)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func intToPgtype(value *int) pgtype.Int4 {
	if value == nil {
		return pgtype.Int4{}
	}

	return pgtype.Int4{
		Int32: int32(*value),
		Valid: true,
	}
}
