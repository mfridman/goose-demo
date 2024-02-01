package migrations

import (
	"context"
	"database/sql"

	"github.com/moby/moby/pkg/namesgenerator"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(up00002, nil)
}

func up00002(ctx context.Context, tx *sql.Tx) error {
	// Add 10 random users
	for i := 0; i < 10; i++ {
		username := namesgenerator.GetRandomName(3)
		_, err := tx.ExecContext(ctx, "INSERT INTO users (username) VALUES (?)", username)
		if err != nil {
			return err
		}
	}
	return nil
}
