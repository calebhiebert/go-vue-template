package main

import (
	"net/http"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// GetMe godoc
// @Summary Gets information on the current user
// @Produce json
// @Success 200 {object} models.User
// @Router /users/me [get]
func (*Controller) GetMe(c *gin.Context) {
	user := extractVerifiedUser(c)

	c.JSON(http.StatusOK, user)
}

// ListUsers godoc
// @Summary Gets a list of all users
// @Produce json
// @Success 200 {object} models.UserSlice
// @Router /admin/users [get]
func (*Controller) ListUsers(c *gin.Context) {
	limit, offset := api.ExtractLimitOffset(c)

	users, err := models.Users(qm.Limit(limit), qm.Offset(offset)).AllG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, users)
}