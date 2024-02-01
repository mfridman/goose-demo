package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mfridman/goose-demo/data/sql/migrations"
	"github.com/pressly/goose/v3"
	"github.com/pressly/goose/v3/database"
	_ "modernc.org/sqlite"
)

func main() {
	log.SetFlags(0)
	ctx := context.Background()

	tmp := os.TempDir()
	defer os.RemoveAll(tmp)
	db, err := sql.Open("sqlite", tmp+"/goose.db")
	if err != nil {
		log.Fatal(err)
	}

	provider, err := goose.NewProvider(database.DialectSQLite3, db, migrations.Embed)
	if err != nil {
		log.Fatal(err)
	}
	// List migration sources the provider is aware of.
	fmt.Println("\n=== migration list ===")
	sources := provider.ListSources()
	for _, s := range sources {
		log.Printf("%-3s %-2v %v\n", s.Type, s.Version, filepath.Base(s.Path))
		// sql 1  00001_users_table.sql
		// go  2  00002_add_users.go
		// go  3  00003_count_users.go
	}

	// List status of migrations before applying them.
	stats, err := provider.Status(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n=== migration status ===")
	for _, s := range stats {
		fmt.Printf("%-3s %-2v %v\n", s.Source.Type, s.Source.Version, s.State)
	}

	fmt.Println("\n=== log migration output  ===")
	results, err := provider.Up(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n=== migration results  ===")
	for _, r := range results {
		fmt.Printf("%-3s %-2v done: %v\n", r.Source.Type, r.Source.Version, r.Duration)
	}
}
