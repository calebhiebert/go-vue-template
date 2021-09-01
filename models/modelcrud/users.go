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

// GetUserByID godoc
// @Summary Gets a single User entity by their id
// @Produce json
// @Success 200 {object} User
// @Router /crud/users/:id [get]
func (*GeneratedCrudController) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	User, err := models.Users(qm.Where("id = ?", id)).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, User)
}

// GetUsers godoc
// @Summary Gets a list for all entities of the User type
// @Produce json
// @Success 200 {object} UserSlice
// @Router /crud/users [get]
func (*GeneratedCrudController) GetUsers(c *gin.Context) {
	limit, offset := api.ExtractLimitOffset(c)

	users, err := models.Users(qm.Limit(limit), qm.Offset(offset)).AllG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, users)
}

func (gcc *GeneratedCrudController) RegisterUsers(rg *gin.RouterGroup) {
	rg.GET("/users/:id", gcc.GetUserByID)
	rg.GET("/users", gcc.GetUsers)
}
