package sbi

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"

	"github.com/nycu-ucr/openapi"
	"github.com/nycu-ucr/openapi/models"
	"github.com/nycu-ucr/udm/internal/logger"
)

func (s *Server) getHttpCallBackRoutes() []Route {
	return []Route{
		{
			"Index",
			http.MethodGet,
			"/",
			s.HandleIndex,
		},

		{
			"DataChangeNotificationToNF",
			http.MethodPost,
			"/sdm-subscriptions",
			s.HandleDataChangeNotificationToNF,
		},
	}
}

func (s *Server) HandleDataChangeNotificationToNF(c *gin.Context) {
	var dataChangeNotify models.DataChangeNotify
	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.CallbackLog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	err = openapi.Deserialize(&dataChangeNotify, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.CallbackLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	supi := c.Params.ByName("supi")

	logger.CallbackLog.Infof("Handle DataChangeNotificationToNF")

	s.Processor().DataChangeNotificationProcedure(c, dataChangeNotify.NotifyItems, supi)
}
