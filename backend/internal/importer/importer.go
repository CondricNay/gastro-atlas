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

	for _, event := range ingredient.Events {
		_, err := queries.CreateEvent(
			ctx,
			sqlc.CreateEventParams{
				IngredientID: ingredientID,
				Description:  event.Description,
				TimePeriod:   event.TimePeriod,
				Location:     event.Location,

				StartYear: pgtype.Int4{
					Int32: func() int32 {
						if event.StartYear != nil {
							return *event.StartYear
						}
						return 0
					}(),
					Valid: event.StartYear != nil,
				},

				EndYear: pgtype.Int4{
					Int32: func() int32 {
						if event.EndYear != nil {
							return *event.EndYear
						}
						return 0
					}(),
					Valid: event.EndYear != nil,
				},

				Latitude: pgtype.Float8{
					Float64: func() float64 {
						if event.Latitude != nil {
							return *event.Latitude
						}
						return 0
					}(),
					Valid: event.Latitude != nil,
				},

				Longitude: pgtype.Float8{
					Float64: func() float64 {
						if event.Longitude != nil {
							return *event.Longitude
						}
						return 0
					}(),
					Valid: event.Longitude != nil,
				},
			},
		)

		if err != nil {
			return fmt.Errorf(
				"create event %q: %w",
				event.Description,
				err,
			)
		}

		fmt.Printf("✓ %s\n", event.Description)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}
