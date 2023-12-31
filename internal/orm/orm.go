package orm

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func New(db *sql.DB) *bun.DB {
	client := bun.NewDB(db, pgdialect.New())
	return client
}
