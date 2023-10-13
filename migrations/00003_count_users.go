package migrations

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(up00003, nil)
}

func up00003(ctx context.Context, tx *sql.Tx) error {
	// Get 10 users
	rows, err := tx.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return err
		}
		users = append(users, user)
	}
	if err := rows.Close(); err != nil {
		return err
	}
	if err := rows.Err(); err != nil {
		return err
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
