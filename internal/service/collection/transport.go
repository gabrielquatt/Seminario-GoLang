package collection

import (
	"net/http"

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

	// obtener todos los juegos
	list = append(list, &endpoint{
		method:   "GET",
		path:     "/games",
		function: getAll(s),
	})

	// agrego un juego
	list = append(list, &endpoint{
		method:   "POST",
		path:     "/NewGame",
		function: postGame(s),
	})

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"games": s.GetAll(),
		})
	}
}

func postGame(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.Query("title")
		description := c.Query("description")
		developer := c.Query("developer")
		c.JSON(http.StatusOK, gin.H{
			"player": s.PostGame(title, description, developer),
		})
	}
}

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
