package importer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5"
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

	ingredientID, err := insertIngredient(ctx, tx, ingredient)
	if err != nil {
		return err
	}

	for _, place := range ingredient.Places {
		placeID, err := insertPlace(ctx, tx, place)
		if err != nil {
			return err
		}

		if err := insertIngredientPlace(
			ctx,
			tx,
			ingredientID,
			placeID,
			place,
		); err != nil {
			return err
		}

		fmt.Printf("✓ %s\n", place.Name)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func insertIngredient(
	ctx context.Context,
	tx pgx.Tx,
	ingredient IngredientData,
) (int, error) {
	var id int

	err := tx.QueryRow(
		ctx,
		`
		INSERT INTO ingredients (name, slug, description)
		VALUES ($1, $2, $3)
		ON CONFLICT (slug)
		DO UPDATE SET
			name = EXCLUDED.name,
			description = EXCLUDED.description
		RETURNING id
		`,
		ingredient.Name,
		ingredient.Slug,
		ingredient.Description,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("insert ingredient: %w", err)
	}

	return id, nil
}

func insertPlace(
	ctx context.Context,
	tx pgx.Tx,
	place PlaceData,
) (int, error) {
	var id int

	err := tx.QueryRow(
		ctx,
		`
		INSERT INTO places (
			name,
			type,
			latitude,
			longitude
		)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (name)
		DO UPDATE SET
			type = EXCLUDED.type,
			latitude = EXCLUDED.latitude,
			longitude = EXCLUDED.longitude
		RETURNING id
		`,
		place.Name,
		place.Type,
		place.Latitude,
		place.Longitude,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf(
			"insert place %q: %w",
			place.Name,
			err,
		)
	}

	return id, nil
}

func insertIngredientPlace(
	ctx context.Context,
	tx pgx.Tx,
	ingredientID int,
	placeID int,
	place PlaceData,
) error {
	_, err := tx.Exec(
		ctx,
		`
		INSERT INTO ingredient_places (
			ingredient_id,
			place_id,
			relationship,
			start_year,
			end_year,
			notes
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (ingredient_id, place_id)
		DO UPDATE SET
			relationship = EXCLUDED.relationship,
			start_year = EXCLUDED.start_year,
			end_year = EXCLUDED.end_year,
			notes = EXCLUDED.notes
		`,
		ingredientID,
		placeID,
		place.Relationship,
		place.StartYear,
		place.EndYear,
		place.Notes,
	)

	if err != nil {
		return fmt.Errorf(
			"insert relationship for %q: %w",
			place.Name,
			err,
		)
	}

	return nil
}