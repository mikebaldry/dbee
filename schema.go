package dbee

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Schema manages the database schema
type Schema struct {
	db *sqlx.DB
}

func (s *Schema) ensureMigrationTableExists() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS migrations (version integer UNIQUE NOT NULL)`)

	return err
}

// DatabaseVersion returns the current version on the database
func (s *Schema) DatabaseVersion() (int, error) {
	var version sql.NullInt64

	if err := s.db.Get(&version, `SELECT MAX(version) FROM migrations`); err != nil {
		return 0, errors.Wrap(err, "unable to get latest migration version")
	}

	if !version.Valid {
		return 0, errors.New("latest migration version invalid")
	}

	return int(version.Int64), nil
}
