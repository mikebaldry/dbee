package dbee

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// DB represents a connection to a database
type DB struct {
	db     *sqlx.DB
	Schema *Schema
}

// NewDB returns a new DB from the passed Configuration
func NewDB(config Configuration) (*DB, error) {
	db, err := sqlx.Connect(config.Driver, config.ConnectionString)

	if err != nil {
		return nil, errors.Wrap(err, "sqlx.Connect returned an error")
	}

	result := DB{
		db: db,
		Schema: &Schema{
			db: db,
		},
	}

	return &result, nil
}
