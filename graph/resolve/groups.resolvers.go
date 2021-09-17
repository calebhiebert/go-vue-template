package resolve

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/graph/model"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/calebhiebert/go-vue-template/models/gql"
	"github.com/calebhiebert/go-vue-template/util"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (r *mutationResolver) CreateGroup(ctx context.Context, c *model.CreateGroup) (*models.Group, error) {
	var group *models.Group

	// Get the user making this request
	currentUser := util.ExtractUser(ctx)

	if currentUser == nil {
		return nil, api.MissingUserError
	}

	// Perform all operations in a db transaction
	err := util.DoTX(ctx, func(ctx context.Context, tx *sql.Tx) error {

		// Use the generated create group helper to create a group
		g, err := gql.CreateGroupTX(ctx, tx, c, func(group *models.Group, currentUser *models.User) error {
			if currentUser == nil {
				return api.MissingUserError
			}

			group.CreatedByID = currentUser.ID

			return nil
		})
		if err != nil {
			return err
		}

		group = g

		// Go through submitted invites and either add the users directly, or create invites for them
		for _, invite := range c.Invites {
			_, err := util.DoInvite(ctx, tx, g, nil, invite, currentUser)
			if err != nil {
				return nil
			}
		}

		groupUser := models.GroupUser{
			GroupID: g.ID,
			UserID:  g.CreatedByID,
		}

		gucIErr := groupUser.Upsert(ctx, tx, true, []string{"group_id", "user_id"}, boil.Infer(), boil.Infer())
		if gucIErr != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return group, nil
}
