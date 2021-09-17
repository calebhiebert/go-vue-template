package resolve

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/graph/model"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/calebhiebert/go-vue-template/util"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *mutationResolver) AcceptInvite(ctx context.Context, shortcode string) (*model.InviteAcceptResult, error) {
	currentUser := util.ExtractUser(ctx)

	if currentUser == nil {
		return nil, api.MissingUserError
	}

	invite, err := models.Invites(
		qm.Where("i.id = ?", shortcode),
		qm.InnerJoin("invite_shortcodes i on invites.id = i.invite_id"),
	).OneG(ctx)
	if err != nil {
		return nil, err
	}

	var res model.InviteAcceptResult

	err = util.DoTX(ctx, func(ctx context.Context, tx *sql.Tx) error {
		if invite.EventID.Valid {
			eventUser := models.EventUser{
				UserID:   currentUser.ID,
				EventID:  invite.EventID.String,
				IsOwner:  false,
				Accepted: true,
			}

			err := eventUser.Insert(ctx, tx, boil.Infer())
			if err != nil {
				return err
			}

			res.EventID = &invite.EventID.String
		} else if invite.GroupID.Valid {
			groupUser := models.GroupUser{
				UserID:  currentUser.ID,
				GroupID: invite.EventID.String,
			}

			err := groupUser.Insert(ctx, tx, boil.Infer())
			if err != nil {
				return err
			}

			res.GroupID = &invite.GroupID.String
		}

		_, err = models.InviteShortcodes(qm.Where("id = ?", shortcode)).DeleteAll(ctx, tx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &res, nil
}
