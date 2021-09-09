// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package modelcrud

import (
	"net/http"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type APIUpdateTokenIssuanceRequest struct {
	UserID    *string `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	IPAddress *string `boil:"ip_address" json:"ip_address" toml:"ip_address" yaml:"ip_address"`
}

type UpdateTokenIssuanceRequest struct {
	UserID    *string `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id"`
	IPAddress *string `boil:"ip_address" json:"ip_address,omitempty" toml:"ip_address" yaml:"ip_address"`
}

// UpdateTokenIssuanceByID godoc
// @Summary Updates a single TokenIssuance entity based on their id
// @Produce json
// @Accept json
// @Param req body APIUpdateTokenIssuanceRequest true "Update parameters"
// @Param id path string true "TokenIssuance id"
// @Success 200 {object} APITokenIssuance
// @Router /crud/tokenIssuances/:id [put]
func (*GeneratedCrudController) UpdateTokenIssuanceByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	var updateReq UpdateTokenIssuanceRequest

	err := c.BindJSON(&updateReq)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	existingTokenIssuance, err := models.TokenIssuances(qm.Where("id = ?", id), qm.For("UPDATE")).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	if updateReq.UserID != nil {
		existingTokenIssuance.UserID = *updateReq.UserID
	}

	if updateReq.IPAddress != nil {
		existingTokenIssuance.IPAddress = *updateReq.IPAddress
	}

	_, err = existingTokenIssuance.UpdateG(c.Request.Context(), boil.Infer())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, existingTokenIssuance)
}
