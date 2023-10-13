package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/mfridman/goose-demo/migrations"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

func main() {
	log.SetFlags(0)

	tmp := os.TempDir()
	defer os.RemoveAll(tmp)
	db, err := sql.Open("sqlite", tmp+"/goose.db")
	if err != nil {
		log.Fatal(err)
	}
	goose.SetDialect("sqlite3")
	goose.SetBaseFS(migrations.Embed)

	if err := goose.Up(db, "."); err != nil {
		panic(err)
	}
	if err := goose.Version(db, "."); err != nil {
		log.Fatal(err)
	}
}
