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
		v, _ := strconv.Atoi(aux)
		i := int64(v)

		err := s.DeleteGame(i)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}

	}
}

func deleteAllGame(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		err := s.DeleteAllGames()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
	}
}

func postGame(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		title := c.Query("Title")
		description := c.Query("Description")
		developer := c.Query("Developer")

		g, err := s.PostGame(NewGame(0, title, description, developer))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":   "OK",
				"New Game": g,
			})
		}
	}
}

func editGame(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		aux := c.Param("ID")
		v, _ := strconv.Atoi(aux)
		i := int64(v)

		title := c.Query("Title")
		description := c.Query("Description")
		developer := c.Query("Developer")

		g, err := s.EditGame(NewGame(i, title, description, developer))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":    "OK",
				"Game Edit": g,
			})
		}
	}
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		g, err := s.GetAll()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
				"Games":  g,
			})
		}
	}
}

func getGameById(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		aux := c.Param("ID")
		v, _ := strconv.Atoi(aux)
		i := int64(v)

		g, err := s.GetGameById(i)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
				"Game":   g,
			})
		}
	}
}

//-------------------------------------------------------------------//

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
