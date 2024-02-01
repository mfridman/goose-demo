package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mfridman/goose-demo/gen/dbstore"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(up00003, nil)
}

func up00003(ctx context.Context, tx *sql.Tx) error {
	q := dbstore.New(tx)
	users, err := q.ListUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to list users: %w", err)
	}
	// Check that we have 10 users
	if len(users) != 10 {
		return fmt.Errorf("expected 10 users, got %d", len(users))
	}
	for _, user := range users {
		fmt.Printf("User: %d, %s\n", user.ID, user.Username)
	}
	return nil
}

type User struct {
	ID       int
	Username string
}
