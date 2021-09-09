package graph

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/calebhiebert/go-vue-template/models"
)

func sanitizeLimitOffset(lim *int, off *int) (limit int, offset int) {
	if lim == nil {
		limit = 10
	} else if *lim <= 0 {
		limit = 10
	} else if *lim > 100 {
		limit = 100
	} else {
		limit = *lim
	}

	if off == nil {
		offset = 0
	} else if *off <= 0 {
		offset = 0
	} else {
		offset = *off
	}

	return
}

func extractUser(ctx context.Context) *models.User {
	raw, ok := ctx.Value("user").(*models.User)
	if !ok {
		return nil
	}

	return raw
}

func DoesPreload(ctx context.Context, field string) bool {
	preloads := GetPreloads(ctx)

	fmt.Println("PRELOADS", field, preloads)

	for _, pl := range preloads {
		if pl == field {
			return true
		}
	}

	return false
}

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}

