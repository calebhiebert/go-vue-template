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

type APIUpdateEventUserRequest struct {
	UserID   *string `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	EventID  *string `boil:"event_id" json:"event_id" toml:"event_id" yaml:"event_id"`
	IsOwner  *bool   `boil:"is_owner" json:"is_owner" toml:"is_owner" yaml:"is_owner"`
	Accepted *bool   `boil:"accepted" json:"accepted" toml:"accepted" yaml:"accepted"`
}

type UpdateEventUserRequest struct {
	UserID   *string `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id"`
	EventID  *string `boil:"event_id" json:"event_id,omitempty" toml:"event_id" yaml:"event_id"`
	IsOwner  *bool   `boil:"is_owner" json:"is_owner,omitempty" toml:"is_owner" yaml:"is_owner"`
	Accepted *bool   `boil:"accepted" json:"accepted,omitempty" toml:"accepted" yaml:"accepted"`
}

// UpdateEventUserByID godoc
// @Summary Updates a single EventUser entity based on their id
// @Produce json
// @Accept json
// @Param req body APIUpdateEventUserRequest true "Update parameters"
// @Param id path string true "EventUser id"
// @Success 200 {object} APIEventUser
// @Router /crud/eventUsers/:id [put]
func (*GeneratedCrudController) UpdateEventUserByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	var updateReq UpdateEventUserRequest

	err := c.BindJSON(&updateReq)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	existingEventUser, err := models.EventUsers(qm.Where("id = ?", id), qm.For("UPDATE")).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}
	if updateReq.UserID != nil {
		existingEventUser.UserID = *updateReq.UserID
	}

	if updateReq.EventID != nil {
		existingEventUser.EventID = *updateReq.EventID
	}

	if updateReq.IsOwner != nil {
		existingEventUser.IsOwner = *updateReq.IsOwner
	}

	if updateReq.Accepted != nil {
		existingEventUser.Accepted = *updateReq.Accepted
	}

	_, err = existingEventUser.UpdateG(c.Request.Context(), boil.Infer())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, existingEventUser)
}
