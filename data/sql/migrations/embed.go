package migrations

import "embed"

//go:embed *.sql
var Embed embed.FS
