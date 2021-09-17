package util

import (
	"context"
	"github.com/calebhiebert/go-vue-template/models"
)

func ExtractUser(ctx context.Context) *models.User {
	raw, ok := ctx.Value("user").(*models.User)
	if !ok {
		return nil
	}

	return raw
}

