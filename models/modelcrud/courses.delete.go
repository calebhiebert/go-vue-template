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

// DeleteCourseByID godoc
// @Summary Soft deletes a single Course entity based on their id
// @Produce json
// @Success 200 {object} APICourse
// @Param id path string true "Course id"
// @Param hardDelete query string false "Hard delete course"
// @Router /crud/courses/:id [delete]
func (*GeneratedCrudController) DeleteCourseByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	hardDelete := c.Query("hardDelete") == "true"

	existingCourse, err := models.Courses(qm.Where("id = ?", id)).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	_, err = existingCourse.DeleteG(c.Request.Context(), hardDelete)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, existingCourse)
}

// BulkDeleteCoursesByIDs godoc
// @Summary Soft deletes a range of courses by their ids
// @Produce json
// @Success 200 {object} DeletedCount
// @Param req body IDList true "List of ids to delete"
// @Param hardDelete query string false "Hard delete course"
// @Router /crud/courses [delete]
func (*GeneratedCrudController) BulkDeleteCoursesByIDs(c *gin.Context) {

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

	deleted, err := models.Courses(qm.WhereIn("id IN ?", idInterface...)).DeleteAllG(c.Request.Context(), hardDelete)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, DeletedCount{DeletedCount: int(deleted)})
}

// UnDeleteCourseByID godoc
// @Summary Undeletes a course by id
// @Produce json
// @Success 200 {object} APICourse
// @Param id path string true "Course id"
// @Router /crud/courses/:id/unDelete [post]
func (*GeneratedCrudController) UnDeleteCourseByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	deletedCourse, err := models.Courses(qm.Where("id = ?", id), qm.WithDeleted()).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	deletedCourse.DeletedAt = null.Time{
		Valid: false,
	}

	_, err = deletedCourse.UpdateG(c.Request.Context(), boil.Whitelist("deleted_at"))
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, deletedCourse)
}
