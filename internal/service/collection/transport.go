package collection

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

// NewHTTPTransport ...
func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	// GET List All Game
	list = append(list, &endpoint{
		method:   "GET",
		path:     "/games/AllGames",
		function: getAll(s),
	})

	// GET Game By ID
	list = append(list, &endpoint{
		method:   "GET",
		path:     "/games/GetGame/:ID",
		function: getGameById(s),
	})

	// Add New Game
	list = append(list, &endpoint{
		method:   "POST",
		path:     "/games/NewGame",
		function: postGame(s),
	})

	// DELETE ALL GAMES
	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/games/DeleteAllGames",
		function: deleteAllGame(s),
	})

	// DELETE GAME BY ID
	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/games/DeleteGame/:ID",
		function: deleteGame(s),
	})

	// EDIT GAME
	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/games/EditGame/:ID",
		function: editGame(s),
	})

	return list
}

//-------------------------------------------------------------------//

func deleteGame(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		aux := c.Param("ID")
		i, _ := strconv.Atoi(aux)
		c.JSON(http.StatusOK, gin.H{
			"status": s.DeleteGame(i),
		})
	}
}

func deleteAllGame(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": s.DeleteAllGames(),
		})
	}
}

func postGame(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO tratar de capturar valores de un JSON y no por Query

		//tomo por parametro de url los valores a guardar
		title := c.Query("Title")
		description := c.Query("Description")
		developer := c.Query("Developer")
		//luego los envio
		c.JSON(http.StatusOK, gin.H{
			"status": s.PostGame(title, description, developer),
		})
	}
}

func editGame(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		aux := c.Param("ID")
		i, _ := strconv.Atoi(aux)
		//--------------------------------//
		title := c.Query("Title")
		description := c.Query("Description")
		developer := c.Query("Developer")
		//--------------------------------//
		c.JSON(http.StatusOK, gin.H{
			"status": s.EditGame(title, description, developer, i),
		})
	}
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": s.GetAll(),
		})
	}
}

func getGameById(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		aux := c.Param("ID")
		i, _ := strconv.Atoi(aux)

		c.JSON(http.StatusOK, gin.H{
			"status": s.GetGameById(i),
		})
	}
}

//-------------------------------------------------------------------//

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
