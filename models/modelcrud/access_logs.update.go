// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package modelcrud

import (
	"net/http"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type APIUpdateAccessLogRequest struct {
	Path *string `boil:"path" json:"path" toml:"path" yaml:"path"`

	RequestBody *string `boil:"request_body" json:"request_body,omitempty" toml:"request_body" yaml:"request_body,omitempty"`

	RequestHeaders map[string]interface{} `boil:"request_headers" json:"request_headers,omitempty" toml:"request_headers" yaml:"request_headers,omitempty"`

	ResponseBody map[string]interface{} `boil:"response_body" json:"response_body" toml:"response_body" yaml:"response_body"`

	ResponseHeaders    map[string]interface{} `boil:"response_headers" json:"response_headers" toml:"response_headers" yaml:"response_headers"`
	ResponseCode       *int                   `boil:"response_code" json:"response_code" toml:"response_code" yaml:"response_code"`
	ProcessingDuration *int                   `boil:"processing_duration" json:"processing_duration" toml:"processing_duration" yaml:"processing_duration"`
	RequestMethod      *string                `boil:"request_method" json:"request_method" toml:"request_method" yaml:"request_method"`

	UserID    *string `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	IPAddress *string `boil:"ip_address" json:"ip_address" toml:"ip_address" yaml:"ip_address"`
}

type UpdateAccessLogRequest struct {
	Path               *string      `boil:"path" json:"path,omitempty" toml:"path" yaml:"path"`
	RequestBody        *null.String `boil:"request_body" json:"request_body,omitempty" toml:"request_body" yaml:"request_body,omitempty"`
	RequestHeaders     *null.JSON   `boil:"request_headers" json:"request_headers,omitempty" toml:"request_headers" yaml:"request_headers,omitempty"`
	ResponseBody       *types.JSON  `boil:"response_body" json:"response_body,omitempty" toml:"response_body" yaml:"response_body"`
	ResponseHeaders    *types.JSON  `boil:"response_headers" json:"response_headers,omitempty" toml:"response_headers" yaml:"response_headers"`
	ResponseCode       *int         `boil:"response_code" json:"response_code,omitempty" toml:"response_code" yaml:"response_code"`
	ProcessingDuration *int         `boil:"processing_duration" json:"processing_duration,omitempty" toml:"processing_duration" yaml:"processing_duration"`
	RequestMethod      *string      `boil:"request_method" json:"request_method,omitempty" toml:"request_method" yaml:"request_method"`
	UserID             *null.String `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	IPAddress          *string      `boil:"ip_address" json:"ip_address,omitempty" toml:"ip_address" yaml:"ip_address"`
}

// UpdateAccessLogByID godoc
// @Summary Updates a single AccessLog entity based on their id
// @Produce json
// @Accept json
// @Param req body APIUpdateAccessLogRequest true "Update parameters"
// @Param id path string true "AccessLog id"
// @Success 200 {object} APIAccessLog
// @Router /crud/accessLogs/:id [put]
func (*GeneratedCrudController) UpdateAccessLogByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	var updateReq UpdateAccessLogRequest

	err := c.BindJSON(&updateReq)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	existingAccessLog, err := models.AccessLogs(qm.Where("id = ?", id), qm.For("UPDATE")).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	if updateReq.Path != nil {
		existingAccessLog.Path = *updateReq.Path
	}

	if updateReq.RequestBody != nil {
		existingAccessLog.RequestBody = *updateReq.RequestBody
	}

	if updateReq.RequestHeaders != nil {
		existingAccessLog.RequestHeaders = *updateReq.RequestHeaders
	}

	if updateReq.ResponseBody != nil {
		existingAccessLog.ResponseBody = *updateReq.ResponseBody
	}

	if updateReq.ResponseHeaders != nil {
		existingAccessLog.ResponseHeaders = *updateReq.ResponseHeaders
	}

	if updateReq.ResponseCode != nil {
		existingAccessLog.ResponseCode = *updateReq.ResponseCode
	}

	if updateReq.ProcessingDuration != nil {
		existingAccessLog.ProcessingDuration = *updateReq.ProcessingDuration
	}

	if updateReq.RequestMethod != nil {
		existingAccessLog.RequestMethod = *updateReq.RequestMethod
	}

	if updateReq.UserID != nil {
		existingAccessLog.UserID = *updateReq.UserID
	}

	if updateReq.IPAddress != nil {
		existingAccessLog.IPAddress = *updateReq.IPAddress
	}

	_, err = existingAccessLog.UpdateG(c.Request.Context(), boil.Infer())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, existingAccessLog)
}
