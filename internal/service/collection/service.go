package collection

import (
	"fmt"

	"github.com/Gabriel-Quattrini/Seminario-GoLang/internal/config"
	"github.com/jmoiron/sqlx"
)

// Service interface
type Service interface {
	GetAll() []*Game
	PostGame(string, string, string) string
}

// Game ...
type Game struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Developer   string `db:"developer"`
}

type service struct {
	db  *sqlx.DB
	cfg *config.Config
}

// NewService ...
func NewService(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

// GetAll ...
func (s service) GetAll() []*Game {
	var list []*Game

	query := "SELECT * FROM game"
	err := s.db.Select(&list, query)
	if err != nil {
		panic(err.Error())
	}
	return list
}

// PostGame ...
func (s service) PostGame(t string, d string, dev string) string {
	query := `INSERT INTO game (title, description, developer) VALUES (?, ?, ?)`

	res := s.db.MustExec(query, t, d, dev)
	LastID, _ := res.LastInsertId()

	return fmt.Sprintf("New Game ID: %d", LastID)
}
