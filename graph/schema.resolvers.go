package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"os"

	"github.com/calebhiebert/go-vue-template/graph/generated"
	"github.com/calebhiebert/go-vue-template/graph/model"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *accessLogResolver) User(ctx context.Context, obj *model.AccessLog) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]*model.User, error) {
	sLim, sOff := sanitizeLimitOffset(limit, offset)

	users, err := models.Users(qm.Limit(sLim), qm.Offset(sOff)).AllG(ctx)
	if err != nil {
		return nil, err
	}

	return r.converter.ConvertUsers(users), nil
}

func (r *tokenIssuanceResolver) User(ctx context.Context, obj *model.TokenIssuance) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) AvatarURL(ctx context.Context, obj *model.User) (string, error) {
	return fmt.Sprintf("%s/avatar/%s", os.Getenv("HOSTED_URL"), obj.ID), nil
}

// AccessLog returns generated.AccessLogResolver implementation.
func (r *Resolver) AccessLog() generated.AccessLogResolver { return &accessLogResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// TokenIssuance returns generated.TokenIssuanceResolver implementation.
func (r *Resolver) TokenIssuance() generated.TokenIssuanceResolver { return &tokenIssuanceResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type accessLogResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tokenIssuanceResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
