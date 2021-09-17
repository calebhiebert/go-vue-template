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
)

// DeleteGroupUserByID godoc
// @Summary Soft deletes a single GroupUser entity based on their id
// @Produce json
// @Success 200 {object} APIGroupUser
// @Param id path string true "GroupUser id"
// @Param hardDelete query string false "Hard delete groupUser"
// @Router /crud/groupUsers/:id [delete]
func (*GeneratedCrudController) DeleteGroupUserByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	hardDelete := c.Query("hardDelete") == "true"

	existingGroupUser, err := models.GroupUsers(qm.Where("id = ?", id)).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	_, err = existingGroupUser.DeleteG(c.Request.Context(), hardDelete)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, existingGroupUser)
}

// BulkDeleteGroupUsersByIDs godoc
// @Summary Soft deletes a range of groupUsers by their ids
// @Produce json
// @Success 200 {object} DeletedCount
// @Param req body IDList true "List of ids to delete"
// @Param hardDelete query string false "Hard delete groupUser"
// @Router /crud/groupUsers [delete]
func (*GeneratedCrudController) BulkDeleteGroupUsersByIDs(c *gin.Context) {

	var ids IDList

	err := c.BindJSON(&ids)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	hardDelete := c.Query("hardDelete") == "true"

	var idInterface []interface{}

	for _, id := range ids.IDs {
		idInterface = append(idInterface, id)
	}

	deleted, err := models.GroupUsers(qm.WhereIn("id IN ?", idInterface...)).DeleteAllG(c.Request.Context(), hardDelete)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, DeletedCount{DeletedCount: int(deleted)})
}

// UnDeleteGroupUserByID godoc
// @Summary Undeletes a groupUser by id
// @Produce json
// @Success 200 {object} APIGroupUser
// @Param id path string true "GroupUser id"
// @Router /crud/groupUsers/:id/unDelete [post]
func (*GeneratedCrudController) UnDeleteGroupUserByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	deletedGroupUser, err := models.GroupUsers(qm.Where("id = ?", id), qm.WithDeleted()).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	deletedGroupUser.DeletedAt = null.Time{
		Valid: false,
	}

	_, err = deletedGroupUser.UpdateG(c.Request.Context(), boil.Whitelist("deleted_at"))
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, deletedGroupUser)
}