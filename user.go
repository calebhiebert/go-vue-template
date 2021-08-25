package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
