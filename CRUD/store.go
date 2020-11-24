package main

import (
	"fmt"
)

// Collection ...
type Collection struct {
	games map[int]*game
}

// game ...
type game struct {
	ID            int
	title         string
	description   string
	desarrollador string
}

// NewCollection ...
func NewCollection() Collection { //devuelve una instancia de la coleccion
	games := make(map[int]*game)
	return Collection{
		games,
	}
}

// Add ...
func (c Collection) Add(l game) {
	c.games[l.ID] = &l
}

// Print ...
func (c Collection) Print() {
	for _, v := range c.games {
		fmt.Printf("[%v]\t titulo %v autor [%v] \n", v.ID, v.title, v.desarrollador)
	}
}

// FindByID ...
func (c Collection) FindByID(ID int) *game {
	// int ID = indice, ID cada vez que se usa que sea en Mayuscula
	return c.games[ID]
}

// Delete ...
func (c Collection) Delete(ID int) {
	//existe la funcion delete,
	delete(c.games, ID)
	//se le pasa el map y el ID
}

// Update ...
func (c Collection) Update(l game) {
	c.games[l.ID] = &l //c.games[ sub pocision de ID] Es igual a la dirrecion de memoria de l
}
