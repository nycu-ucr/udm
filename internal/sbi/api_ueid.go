package sbi

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"
)

func (s *Server) getUEIDRoutes() []Route {
	return []Route{
		{
			"Index",
			http.MethodGet,
			"/",
			s.HandleIndex,
		},

		{
			"Deconceal",
			http.MethodPost,
			"/deconceal",
			s.HandleDeconceal,
		},
	}
}

func (s *Server) HandleDeconceal(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}
