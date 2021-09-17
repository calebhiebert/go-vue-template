package resolve

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/calebhiebert/go-vue-template/graph/generated"
	"github.com/calebhiebert/go-vue-template/graph/model"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/calebhiebert/go-vue-template/models/gql"
	"github.com/calebhiebert/go-vue-template/util"
)

func (r *mutationResolver) UpdateUserByID(ctx context.Context, id string, u model.UpdateUser) (*models.User, error) {
	return gql.UpdateUserHelper(ctx, id, &u, func(user *models.User, currentUser *models.User) bool {
		return user.ID == currentUser.ID || util.StringSliceContains(currentUser.Roles, "admin")
	})
}

func (r *userResolver) Roles(ctx context.Context, obj *models.User) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
