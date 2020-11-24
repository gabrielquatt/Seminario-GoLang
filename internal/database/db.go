package database

import (
	"errors"

	"github.com/Gabriel-Quattrini/Seminario-GoLang/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// NewDatabase ...
func NewDatabase(cfg *config.Config) (*sqlx.DB, error) {
	switch cfg.DB.Type {
	case "sqlite3":
		db, err := sqlx.Open(cfg.DB.Driver, cfg.DB.Conn)
		if err != nil {
			return nil, err
		}

		err = db.Ping()
		if err != nil {
			return nil, err
		}

		return db, nil
	default:
		return nil, errors.New("invalid db type")
	}
}
