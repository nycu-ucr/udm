package sbi

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"
)

func (s *Server) getReportSMDeliveryStatusRoutes() []Route {
	return []Route{
		{
			"Index",
			http.MethodGet,
			"/",
			s.HandleIndex,
		},

		{
			"ReportSMDeliveryStatus",
			http.MethodPost,
			"/:ueIdentity/sm-delivery-status",
			s.HandleReportSMDeliveryStatus,
		},
	}
}

func (s *Server) HandleReportSMDeliveryStatus(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}
