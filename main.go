package main

import (
	"fmt"
)

func main() {

	//CRUD
	c := NewCollection()

	c.Add(Libro{0, "Libroq", "description", "author", 200}) //Create

	c.Add(Libro{1, "Libro2", "bbbb", "author2", 2000}) //Create

	c.Add(Libro{2, "Libro3", "aaaa", "author2", 900}) //Create

	c.Print() //imprime el map

	l0 := c.FindByID(0) //retorno de un elemento del map
	if l0 != nil {
		fmt.Println("se encontro el ID=0")
	}

	c.Delete(0) //elimina segun el id mandado
	c.Print()
	fmt.Println(".......................")
	c.Update(Libro{2, "----", "----", "----", 203})
	c.Print()

}
