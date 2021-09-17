// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package modelcrud

import (
	"net/http"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// DeleteProfileQuestionByID godoc
// @Summary Soft deletes a single ProfileQuestion entity based on their id
// @Produce json
// @Success 200 {object} APIProfileQuestion
// @Param id path string true "ProfileQuestion id"
// @Router /crud/profileQuestions/:id [delete]
func (*GeneratedCrudController) DeleteProfileQuestionByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	existingProfileQuestion, err := models.ProfileQuestions(qm.Where("id = ?", id)).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	_, err = existingProfileQuestion.DeleteG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, existingProfileQuestion)
}

// BulkDeleteProfileQuestionsByIDs godoc
// @Summary Soft deletes a range of profileQuestions by their ids
// @Produce json
// @Success 200 {object} DeletedCount
// @Param req body IDList true "List of ids to delete"
// @Param hardDelete query string false "Hard delete profileQuestion"
// @Router /crud/profileQuestions [delete]
func (*GeneratedCrudController) BulkDeleteProfileQuestionsByIDs(c *gin.Context) {

	var ids IDList

	err := c.BindJSON(&ids)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	var idInterface []interface{}

	for _, id := range ids.IDs {
		idInterface = append(idInterface, id)
	}

	deleted, err := models.ProfileQuestions(qm.WhereIn("id IN ?", idInterface...)).DeleteAllG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, DeletedCount{DeletedCount: int(deleted)})
}