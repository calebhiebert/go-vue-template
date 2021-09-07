package main

import (
	"context"
	"net/http"

	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// GetMe godoc
// @Summary Gets information on the current user
// @Produce json
// @Success 200 {object} modelcrud.APIUser
// @Router /users/me [get]
func (*Controller) GetMe(c *gin.Context) {
	user := extractVerifiedUser(c)

	c.JSON(http.StatusOK, user)
}

func addUserHooks() {
	models.AddUserHook(boil.BeforeUpdateHook, BeforeUserUpdate)
}

func BeforeUserUpdate(ctx context.Context, executor boil.ContextExecutor, user *models.User) error {
	// Validate user permissions
	err := ACAdminOr(ctx, func(u *models.User) (bool, error) {
		return u.ID == user.ID, nil
	})
	if err != nil {
		return err
	}

	return nil
}