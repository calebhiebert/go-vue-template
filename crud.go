package main

import (
	"net/http"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type GeneratedCrudController struct{}

func RegisterGeneratedCrud(rg *gin.RouterGroup) {
	gcc := GeneratedCrudController{}

	gcc.RegisterUsers(rg)
}

// GetUserByID godoc
// @Summary Gets a single User entity by their id
// @Produce json
// @Success 200 {object} User
// @Router /modelcrud/users/:id [get]
func (*GeneratedCrudController) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	user, err := models.Users(qm.Where("id = ?", id)).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUsers godoc
// @Summary Gets a list for all entities of the User type
// @Produce json
// @Success 200 {object} UserSlice
// @Router /modelcrud/users [get]
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