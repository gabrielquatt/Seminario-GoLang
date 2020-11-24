package main

import (
	"fmt"
)

//--- go run ./CRUD/ *.go ---//

func main() {

	//CRUD
	c := NewCollection()

	c.Add(game{0, "game", "description", "author"}) //Create

	c.Add(game{1, "game2", "bbbb", "author2"}) //Create

	c.Add(game{2, "game3", "aaaa", "author3"}) //Create

	c.Print() //imprime el map

	l0 := c.FindByID(0) //retorno de un elemento del map
	if l0 != nil {
		fmt.Println("se encontro el ID=0")
	}

	c.Delete(0) //elimina segun el id mandado
	c.Print()
	fmt.Println(".......................")
	c.Update(game{2, "----", "----", "----"})
	c.Print()

}
