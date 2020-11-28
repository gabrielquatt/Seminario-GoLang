package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"github.com/Gabriel-Quattrini/Seminario-GoLang/internal/config"
	"github.com/Gabriel-Quattrini/Seminario-GoLang/internal/database"
	"github.com/Gabriel-Quattrini/Seminario-GoLang/internal/service/collection"
)

//---go run cmd/collectionSRV.go -config ./config/config.yaml--//

func main() {

	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	cfg := config.LoadConfig(*configFile)
	db, err := database.NewDatabase(cfg)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	err = createSchema(db) // db de prueba
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := collection.NewService(db, cfg)
	//carga de datos de prueba

	/* s := service.PostGame("DOOM", "juego cargado de prueba", "id sofware")
	fmt.Println(s) */

	/* service.PostGame("DOOM", "juego cargado de prueba", "id sofware")
	service.PostGame("Call Of Duty", "Segundo juego cargado de prueba", "pum")
	service.PostGame("PES 2020", "tercer juego cargado de prueba", "copy paste") */

	httpService := collection.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()

}

func createSchema(db *sqlx.DB) error {
	//creo esquema de la base de datos
	schema := `CREATE TABLE IF NOT EXISTS game ( 
		id integer primary key autoincrement,
		title varchar (100),
		description varchar (100),
		developer varchar (100));`
	// execute a query on the server
	_, err := db.Exec(schema)

	if err != nil {
		panic(err.Error)
	}

	return nil
}
