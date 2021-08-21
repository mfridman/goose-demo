package main

import (
	"database/sql"
	"embed"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	log.SetFlags(0)
	db, err := sql.Open("sqlite3", "embed_example.sql")
	if err != nil {
		log.Fatal(err)
	}
	goose.SetDialect("sqlite3")
	goose.SetBaseFS(embedMigrations)

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
	if err := goose.Version(db, "migrations"); err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	var user struct {
		ID       int
		Username string
	}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			log.Fatal(err)
		}
		log.Println(user.ID, user.Username)
	}
}
