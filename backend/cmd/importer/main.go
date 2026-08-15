package main

import (
	"context"
	"fmt"

	"github.com/CondricNay/gastro-atlas/internal/database"
	"github.com/CondricNay/gastro-atlas/internal/importer"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("load .env: %w", err))
	}

	ctx := context.Background()

	db, err := database.New()
	if err != nil {
		panic(fmt.Errorf("connect to database: %w", err))
	}
	defer db.Close()

	if err := importer.Run(ctx, db, "data/ingredients"); err != nil {
		panic(err)
	}
}
