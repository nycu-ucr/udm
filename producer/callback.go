package producer

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/http_wrapper"
	"github.com/nycu-ucr/openapi/models"
	"github.com/nycu-ucr/udm/logger"
	"github.com/nycu-ucr/udm/producer/callback"
)

// HandleDataChangeNotificationToNFRequest ... Send Data Change Notification
func HandleDataChangeNotificationToNFRequest(request *http_wrapper.Request) *http_wrapper.Response {
	// step 1: log
	logger.CallbackLog.Infof("Handle DataChangeNotificationToNF")

	// step 2: retrieve request
	dataChangeNotify := request.Body.(models.DataChangeNotify)
	supi := request.Params["supi"]

	problemDetails := callback.DataChangeNotificationProcedure(dataChangeNotify.NotifyItems, supi)

	// step 4: process the return value from step 3
	if problemDetails != nil {
		return http_wrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {
		return http_wrapper.NewResponse(http.StatusNoContent, nil, nil)
	}
}
