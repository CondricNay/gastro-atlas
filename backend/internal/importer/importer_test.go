package importer

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func TestImportIngredient(t *testing.T) {
	ctx := context.Background()

	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working directory: %v", err)
	}

	schemaPath := filepath.Join(
		workingDir,
		"..",
		"database",
		"schema.sql",
	)

	container, err := postgres.Run(
		ctx,
		"postgres:17",
		postgres.WithDatabase("gastro_test"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		postgres.WithInitScripts(schemaPath),
		postgres.BasicWaitStrategies(),
	)
	if err != nil {
		t.Fatalf("start postgres container: %v", err)
	}

	t.Cleanup(func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			t.Fatalf("terminate postgres container: %v", err)
		}
	})

	connString, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("get connection string: %v", err)
	}

	db, err := pgxpool.New(ctx, connString)
	if err != nil {
		t.Fatalf("connect to postgres: %v", err)
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		t.Fatalf("ping postgres: %v", err)
	}

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
}