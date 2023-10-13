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
OK   00001_users_table.sql (670.88µs)
OK   00002_add_users.go (404.92µs)
User: 1, intelligent_dubinsky2
User: 2, busy_galileo7
User: 3, confident_feistel1
User: 4, keen_maxwell3
User: 5, objective_burnell6
User: 6, sleepy_jennings2
User: 7, dazzling_dijkstra2
User: 8, naughty_roentgen8
User: 9, clever_wilbur9
User: 10, pedantic_satoshi2
OK   00003_count_users.go (365.21µs)
goose: successfully migrated database to version: 3
goose: version 3
```
