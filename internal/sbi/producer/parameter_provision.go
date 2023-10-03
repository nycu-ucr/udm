package producer

import (
	"context"
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/openapi"
	"github.com/nycu-ucr/openapi/models"
	"github.com/free5gc/udm/internal/logger"
	"github.com/nycu-ucr/util/httpwrapper"
)

func HandleUpdateRequest(request *httpwrapper.Request) *httpwrapper.Response {
	// step 1: log
	logger.PpLog.Infoln("Handle UpdateRequest")

	// step 2: retrieve request
	updateRequest := request.Body.(models.PpData)
	gpsi := request.Params["gpsi"]

	// step 3: handle the message
	problemDetails := UpdateProcedure(updateRequest, gpsi)

	// step 4: process the return value from step 3
	if problemDetails != nil {
		return httpwrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {
		return httpwrapper.NewResponse(http.StatusNoContent, nil, nil)
	}
}

func UpdateProcedure(updateRequest models.PpData, gpsi string) (problemDetails *models.ProblemDetails) {
	clientAPI, err := createUDMClientToUDR(gpsi)
	if err != nil {
		return openapi.ProblemDetailsSystemFailure(err.Error())
	}
	res, err := clientAPI.ProvisionedParameterDataDocumentApi.ModifyPpData(context.Background(), gpsi, nil)
	if err != nil {
		problemDetails = &models.ProblemDetails{
			Status: int32(res.StatusCode),
			Cause:  err.(openapi.GenericOpenAPIError).Model().(models.ProblemDetails).Cause,
			Detail: err.Error(),
		}
		return problemDetails
	}
	defer func() {
		if rspCloseErr := res.Body.Close(); rspCloseErr != nil {
			logger.PpLog.Errorf("ModifyPpData response body cannot close: %+v", rspCloseErr)
		}
	}()
	return nil
}
