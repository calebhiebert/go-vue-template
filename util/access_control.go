package util

import (
	"context"
	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"net/http"
)

type ValidateUserPermissionsFunction func(u *models.User) (bool, error)

func ACAdminOr(ctx context.Context, f ValidateUserPermissionsFunction) error {
	user := ExtractUser(ctx)

	if user == nil {
		return api.NewAPIError("missing-permissions", http.StatusForbidden, "Missing required permissions")
	}

	if StringSliceContains(user.Roles, "admin") {
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
