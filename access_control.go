package main

import (
	"context"
	"net/http"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
)

type ValidateUserPermissionsFunction func(u *models.User) (bool, error)

func ACAdminOr(ctx context.Context, f ValidateUserPermissionsFunction) error {
	user := extractUser(ctx)

	if user == nil {
		return api.NewAPIError("missing-permissions", http.StatusForbidden, "Missing required permissions")
	}

	if stringSliceContains(user.Roles, "admin") {
		return nil
	}

	valid, err := f(user)
	if err != nil {
		return err
	}

	if !valid {
		return api.NewAPIError("missing-permissions", http.StatusForbidden, "Missing required permissions")
	}

	return nil
}