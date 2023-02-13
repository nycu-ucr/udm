package util

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/openapi/models"
)

func ProblemDetailsSystemFailure(detail string) *models.ProblemDetails {
	return &models.ProblemDetails{
		Title:  "System failure",
		Status: http.StatusInternalServerError,
		Detail: detail,
		Cause:  "SYSTEM_FAILURE",
	}
}
