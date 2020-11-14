package main

import (
	"fmt"
)

//al ser publicos me pide comentarlos

// Collection ...
type Collection struct {
	libros map[int]*Libro
}

// Libro ...
type Libro struct {
	ID          int
	title       string
	description string
	author      string
	pages       int
}

// NewCollection ...
func NewCollection() Collection { //devuelve una instancia de la coleccion
	libros := make(map[int]*Libro)
	return Collection{
		libros,
	}
}

// Add ...
func (c Collection) Add(l Libro) {
	c.libros[l.ID] = &l
}

// Print ...
func (c Collection) Print() {
	for _, v := range c.libros {
		fmt.Printf("[%v]\t titulo %v autor [%v]  paginas: [%v] \n", v.ID, v.title, v.author, v.pages)
	}
}

// FindByID ...
func (c Collection) FindByID(ID int) *Libro { // int ID = indice, ID cada vez que se usa que sea en Mayuscula
	return c.libros[ID]
}

// Delete ...
func (c Collection) Delete(ID int) {
	//existe la funcion delete,
	delete(c.libros, ID)
	//se le pasa el map y el ID
}

// Update ...
func (c Collection) Update(l Libro) {
	c.libros[l.ID] = &l //c.libros[ sub pocision de ID] Es igual a la dirrecion de memoria de l
}
