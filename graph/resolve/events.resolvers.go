package resolve

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"fmt"
	"math"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/graph/model"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/calebhiebert/go-vue-template/models/gql"
	"github.com/calebhiebert/go-vue-template/util"
	"github.com/gofrs/uuid"
	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (r *mutationResolver) CreateEvent(ctx context.Context, e model.CreateEvent) (*models.EventUser, error) {
	var eu *models.EventUser

	// Complete the creation inside a transaction, because it contains multiple operations
	err := util.DoTX(ctx, func(ctx context.Context, tx *sql.Tx) error {
		// Create the event
		ev, err := gql.CreateEventTX(ctx, tx, &e, func(event *models.Event, currentUser *models.User) error {
			if currentUser == nil {
				return api.MissingUserError
			}

			event.CreatedByID = currentUser.ID

			return nil
		})
		if err != nil {
			return err
		}

		// Create a matching event user for the user who created the event
		eu = &models.EventUser{
			UserID:   ev.CreatedByID,
			EventID:  ev.ID,
			IsOwner:  true,
			Accepted: true,
		}

		err = eu.Insert(ctx, tx, boil.Infer())
		if err != nil {
			return err
		}

		// If the created event has a player count higher than 4, create enough sub-events to satisfy the player count
		if e.PlayerCount > 4 {
			neededEventCount := int(math.Ceil(float64(e.PlayerCount) / 4.0))

			for i := 0; i < neededEventCount; i++ {

				subEvent := models.Event{
					ID:          uuid.Must(uuid.NewV4()).String(),
					GameTypeID:  ev.GameTypeID,
					EventDate:   ev.EventDate,
					PlayerCount: 4,
					CreatedByID: ev.CreatedByID,
					CourseID:    ev.CourseID,
					ParentID:    null.StringFrom(ev.ID),
					Name:        fmt.Sprintf("%s - Group %d", ev.Name, i+1),
				}

				err := subEvent.Insert(ctx, tx, boil.Infer())
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return eu, nil
}

func (r *mutationResolver) InviteToEvent(ctx context.Context, eventID string, invites []*model.DoInvite) ([]string, error) {
	currentUser := util.ExtractUser(ctx)

	if currentUser == nil {
		return nil, api.MissingUserError
	}

	event, err := models.FindEventG(ctx, eventID)
	if err != nil {
		return nil, err
	}

	if event.CreatedByID != currentUser.ID {
		return nil, api.MissingPermissionsError
	}

	for _, invite := range invites {
		_, err := util.DoInvite(ctx, boil.GetContextDB(), nil, event, invite, currentUser)
		if err != nil {
			fmt.Println("Failed to send invite", err.Error())
		}
	}

	return nil, err
}
