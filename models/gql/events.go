// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package gql

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/graph/model"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/calebhiebert/go-vue-template/util"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type CanUpdateEventFunc func(event *models.Event, currentUser *models.User) bool

func UpdateEventHelper(ctx context.Context, id string, ud *model.UpdateEvent, canUpdate CanUpdateEventFunc) (*models.Event, error) {
	existing, err := models.Events(qm.Where("id = ?", id)).OneG(ctx)
	if err != nil {
		return nil, err
	}

	if !canUpdate(existing, util.ExtractUser(ctx)) {
		return nil, api.NewAPIError("missing-permissions", http.StatusForbidden, "Missing required permissions")
	}

	if ud.GameTypeID != nil {

		existing.GameTypeID = *ud.GameTypeID

	}
	if ud.EventDate != nil {

		existing.EventDate = *ud.EventDate

	}
	if ud.PlayerCount != nil {

		existing.PlayerCount = *ud.PlayerCount

	}
	if ud.Name != nil {

		existing.Name = *ud.Name

	}

	if ud.CourseID != nil {

		existing.CourseID = *ud.CourseID

	}
	if ud.ParentID != nil {

		existing.ParentID = *ud.ParentID

	}
	if ud.Public != nil {

		existing.Public = *ud.Public

	}

	_, err = existing.UpdateG(ctx, boil.Infer())
	if err != nil {
		return nil, err
	}

	return existing, nil
}

type PreEventCreateFunc func(event *models.Event, currentUser *models.User) error

func CreateEvent(ctx context.Context, c *model.CreateEvent, pre PreEventCreateFunc) (*models.Event, error) {
	newEvent := models.Event{}

	newEvent.ID = uuid.Must(uuid.NewV4()).String()

	newEvent.GameTypeID = c.GameTypeID

	newEvent.EventDate = c.EventDate

	newEvent.PlayerCount = c.PlayerCount

	if c.Name != nil {

		newEvent.Name = *c.Name

	}

	newEvent.CourseID = c.CourseID

	if c.ParentID != nil {

		newEvent.ParentID = *c.ParentID

	}

	if c.Public != nil {

		newEvent.Public = *c.Public

	}

	if err := pre(&newEvent, util.ExtractUser(ctx)); err != nil {
		return nil, err
	}

	err := newEvent.InsertG(ctx, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &newEvent, nil
}

func CreateEventTX(ctx context.Context, tx *sql.Tx, c *model.CreateEvent, pre PreEventCreateFunc) (*models.Event, error) {
	newEvent := models.Event{}

	newEvent.ID = uuid.Must(uuid.NewV4()).String()

	newEvent.GameTypeID = c.GameTypeID

	newEvent.EventDate = c.EventDate

	newEvent.PlayerCount = c.PlayerCount

	if c.Name != nil {

		newEvent.Name = *c.Name

	}

	newEvent.CourseID = c.CourseID

	if c.ParentID != nil {

		newEvent.ParentID = *c.ParentID

	}

	if c.Public != nil {

		newEvent.Public = *c.Public

	}

	if err := pre(&newEvent, util.ExtractUser(ctx)); err != nil {
		return nil, err
	}

	err := newEvent.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &newEvent, nil
}