# README

This repo was used in a blog post about [pressly/goose](https://github.com/pressly/goose).

Post can be found here:

https://pressly.github.io/goose/blog/2021/embed-sql-migrations/

## Example

To run the example, clone this repository and run:

```
$ go run ./cmd/custom-goose
```

For demonstration purposes we're using a sqlite3 database, but goose supports many other databases.

1. Migration 1 (sql) creates a users table Migration
2. Migration 2 (go) inserts 10 users into the users table
3. Migration 3 (go) counts the number of users and prints them out to the console

The output should be:

```
=== migration list ===
sql 1  00001_users_table.sql
go  2  00002_add_users.go
go  3  00003_count_users.go

=== migration status ===
sql 1  pending
go  2  pending
go  3  pending

=== log migration output  ===
User: 7, affectionate_buck4
User: 5, angry_mclaren3
User: 8, friendly_thompson6
User: 10, funny_babbage8
User: 2, funny_gagarin9
User: 4, infallible_pare9
User: 3, priceless_margulis7
User: 6, silly_lovelace0
User: 1, sleepy_easley0
User: 9, tender_mcclintock0

=== migration results  ===
sql 1  completed in: 763µs
go  2  completed in: 612.417µs
go  3  completed in: 489.042µs
```
