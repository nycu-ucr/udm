/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package subscriberdatamanagement

import (
	"github.com/nycu-ucr/gonet/http"

	"github.com/nycu-ucr/gin"
)

// GetSmsMngData - retrieve a UE's SMS Management Subscription Data
func HTTPGetSmsMngData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
