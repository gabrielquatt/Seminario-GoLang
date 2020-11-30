package collection

import (
	"errors"

	"github.com/Gabriel-Quattrini/Seminario-GoLang/internal/config"
	"github.com/jmoiron/sqlx"
)

// Game ...
type Game struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Developer   string `db:"developer"`
}

type service struct {
	db  *sqlx.DB
	cfg *config.Config
}

// Service interface
type Service interface {
	DeleteGame(int64) error
	DeleteAllGames() error
	PostGame(*Game) (*Game, error)
	EditGame(*Game) (*Game, error)
	GetAll() ([]*Game, error)
	GetGameById(int64) (*Game, error)
}

// NewService ...
func NewService(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

func NewGame(i int64, t string, des string, dev string) *Game {
	return &Game{
		ID:          i,
		Title:       t,
		Description: des,
		Developer:   dev,
	}
}

//----------------------------------------------------------------------------//

// DeleteGame Elimina un juego segun su ID
func (s service) DeleteGame(i int64) error {

	query := `DELETE FROM game WHERE id = ?`
	res, err := s.db.Exec(query, i)

	if err != nil {
		return errors.New("ERROR IN THE DATABASE " + err.Error())
	}

	c, _ := res.RowsAffected()

	if c == 0 {
		return errors.New("GAME NO EXIST")
	}

	return nil
}

// DeleteAllGames Elimina todos los elementos en la base de datos
func (s service) DeleteAllGames() error {

	query := `DELETE FROM game`
	res, err := s.db.Exec(query)

	if err != nil {
		return errors.New("ERROR IN THE DATABASE " + err.Error())
	}

	c, _ := res.RowsAffected()

	if c == 0 {
		return errors.New("NO GAMES ELIMINATED")
	}

	return nil
}

// PostGame Agrega elementos en la base de datos
func (s service) PostGame(g *Game) (*Game, error) {

	query := `INSERT INTO game (title, description, developer) VALUES (?, ?, ?)`
	res, err := s.db.Exec(query, g.Title, g.Description, g.Description)

	if err != nil {
		return nil, errors.New("ERROR IN THE DATABASE " + err.Error())
	}

	i, _ := res.LastInsertId()
	g.ID = i

	return g, nil
}

// EditGame Recive strings, y el ID del Game a editar
func (s service) EditGame(g *Game) (*Game, error) {

	query := `UPDATE game SET title = ?, description = ?, developer = ? WHERE id = ?`
	res, err := s.db.Exec(query, g.Title, g.Description, g.Developer, g.ID)

	if err != nil {
		return nil, errors.New("ERROR IN THE DATABASE " + err.Error())
	}

	c, _ := res.RowsAffected()

	if c == 0 {
		return nil, errors.New("NO GAME EDITED")
	}

	return g, nil
}

// GetAll  Devuelve una lista de todos los elementos en la base de datos
func (s service) GetAll() ([]*Game, error) {

	var list []*Game

	query := `SELECT * FROM game`
	err := s.db.Select(&list, query)

	if err != nil {
		return nil, err
	}

	return list, nil
}

// GetGameById Devuelve un juego segun su ID
func (s service) GetGameById(i int64) (*Game, error) {

	var g Game

	query := `SELECT * FROM game WHERE id = ?`
	err := s.db.Get(&g, query, i)

	if err != nil {
		return nil, err
	}

	return &g, nil
}
