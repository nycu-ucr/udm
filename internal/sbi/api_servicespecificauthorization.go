package sbi

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"
)

func (s *Server) getServiceSpecificAuthorizationRoutes() []Route {
	return []Route{
		{
			"Index",
			http.MethodGet,
			"/",
			s.HandleIndex,
		},

		{
			"ServiceSpecificAuthorization",
			http.MethodPost,
			"/:ueIdentity/:serviceType/authorize",
			s.HandleServiceSpecificAuthorization,
		},

		{
			"ServiceSpecificAuthorizationRemoval",
			http.MethodPost,
			"/:ueIdentity/:serviceType/remove",
			s.HandleServiceSpecificAuthorizationRemoval,
		},
	}
}

func (s *Server) HandleServiceSpecificAuthorization(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}

func (s *Server) HandleServiceSpecificAuthorizationRemoval(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}
