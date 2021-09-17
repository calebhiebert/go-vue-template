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

type APIUpdateGroupUserRequest struct {
	GroupID *string `boil:"group_id" json:"group_id" toml:"group_id" yaml:"group_id"`
	UserID  *string `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
}

type UpdateGroupUserRequest struct {
	GroupID *string `boil:"group_id" json:"group_id,omitempty" toml:"group_id" yaml:"group_id"`
	UserID  *string `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id"`
}

// UpdateGroupUserByID godoc
// @Summary Updates a single GroupUser entity based on their id
// @Produce json
// @Accept json
// @Param req body APIUpdateGroupUserRequest true "Update parameters"
// @Param id path string true "GroupUser id"
// @Success 200 {object} APIGroupUser
// @Router /crud/groupUsers/:id [put]
func (*GeneratedCrudController) UpdateGroupUserByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	var updateReq UpdateGroupUserRequest

	err := c.BindJSON(&updateReq)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	existingGroupUser, err := models.GroupUsers(qm.Where("id = ?", id), qm.For("UPDATE")).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}
	if updateReq.GroupID != nil {
		existingGroupUser.GroupID = *updateReq.GroupID
	}

	if updateReq.UserID != nil {
		existingGroupUser.UserID = *updateReq.UserID
	}

	_, err = existingGroupUser.UpdateG(c.Request.Context(), boil.Infer())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, existingGroupUser)
}
