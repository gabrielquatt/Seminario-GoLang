package collection

import (
	"errors"
	"fmt"

	"github.com/Gabriel-Quattrini/Seminario-GoLang/internal/config"
	"github.com/jmoiron/sqlx"
)

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

// Service interface
type Service interface {
	GetAll() []*Game
	GetGameById(int) *Game
	PostGame(string, string, string) string
	DeleteAllGames() string
	DeleteGame(int) string
	EditGame(string, string, string, int) string
}

// NewService ...
func NewService(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

//----------------------------------------------------------------------------//

// DeleteGame elimina un juego segun su ID
func (s service) DeleteGame(i int) string {
	query := `DELETE FROM game WHERE id = ?`
	res, err := s.db.Exec(query, i)

	if err != nil {
		return fmt.Sprintf("%v", errors.New("ERROR "+err.Error()))
	}
	RowsAffected, _ := res.RowsAffected()

	return fmt.Sprintf("GAME Column Delete: %d", RowsAffected)
}

// DeleteAllGames Elimina todos los elementos en la base de datos
func (s service) DeleteAllGames() string {
	query := `DELETE FROM game`
	_, err := s.db.Exec(query)

	if err != nil {
		return "ERROR, NO SE PUDO ELIMINAR GAME"
	}

	return "Clear Complet"
}

// PostGame Agrega elementos en la base de datos
func (s service) PostGame(t string, d string, dev string) string {
	query := `INSERT INTO game (title, description, developer) VALUES (?, ?, ?)`

	res := s.db.MustExec(query, t, d, dev)

	LastID, _ := res.LastInsertId() //retorna el ultimo ID añadido

	return fmt.Sprintf("New Game ID: %d", LastID)
}

// EditGame recivo strings, y el id del juego a editar
func (s service) EditGame(t string, des string, dev string, i int) string {
	query := `UPDATE game SET title = ?, description = ?, developer = ? WHERE id = ?`
	_, err := s.db.Exec(query, t, des, dev, i)

	if err != nil {
		return fmt.Sprintf("%v", errors.New("ERROR "+err.Error()))
	}

	return fmt.Sprintf("Column Game Edit: %v", i)
}

// GetAll  Devuelve una lista de todos los elementos en la base de datos
func (s service) GetAll() []*Game {
	var list []*Game
	query := `SELECT * FROM game`
	err := s.db.Select(&list, query)

	if err != nil {
		panic(err.Error())
	}
	return list
}

// getGameById Devuelve un juego segun su ID
func (s service) GetGameById(i int) *Game {
	var g Game
	query := `SELECT * FROM game WHERE id = ?`
	err := s.db.Get(&g, query, i)

	if err != nil {
		return nil
	}
	return &g
}

//---------------------------------------------------------------------------------//

//TODO -manejar los errores mejor y no solo devolver un string
// 	   -no ignorar sql:Result cuando uso Ecex en editar y eliminar

// link de donde saque referencias para la realizacion los metodos
// https://parzibyte.me/blog/2018/12/10/crud-golang-mysql/
