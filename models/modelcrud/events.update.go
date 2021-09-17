// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package modelcrud

import (
	"net/http"
	"time"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type APIUpdateEventRequest struct {
	GameTypeID  *string    `boil:"game_type_id" json:"game_type_id" toml:"game_type_id" yaml:"game_type_id"`
	EventDate   *time.Time `boil:"event_date" json:"event_date" toml:"event_date" yaml:"event_date"`
	PlayerCount *int       `boil:"player_count" json:"player_count" toml:"player_count" yaml:"player_count"`
	Name        *string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedByID *string    `boil:"created_by_id" json:"created_by_id" toml:"created_by_id" yaml:"created_by_id"`
	CourseID    *string    `boil:"course_id" json:"course_id" toml:"course_id" yaml:"course_id"`

	ParentID *string `boil:"parent_id" json:"parent_id,omitempty" toml:"parent_id" yaml:"parent_id,omitempty"`
	Public   *bool   `boil:"public" json:"public" toml:"public" yaml:"public"`
}

type UpdateEventRequest struct {
	GameTypeID  *string      `boil:"game_type_id" json:"game_type_id,omitempty" toml:"game_type_id" yaml:"game_type_id"`
	EventDate   *time.Time   `boil:"event_date" json:"event_date,omitempty" toml:"event_date" yaml:"event_date"`
	PlayerCount *int         `boil:"player_count" json:"player_count,omitempty" toml:"player_count" yaml:"player_count"`
	Name        *string      `boil:"name" json:"name,omitempty" toml:"name" yaml:"name"`
	CreatedByID *string      `boil:"created_by_id" json:"created_by_id,omitempty" toml:"created_by_id" yaml:"created_by_id"`
	CourseID    *string      `boil:"course_id" json:"course_id,omitempty" toml:"course_id" yaml:"course_id"`
	ParentID    *null.String `boil:"parent_id" json:"parent_id,omitempty" toml:"parent_id" yaml:"parent_id,omitempty"`
	Public      *bool        `boil:"public" json:"public,omitempty" toml:"public" yaml:"public"`
}

// UpdateEventByID godoc
// @Summary Updates a single Event entity based on their id
// @Produce json
// @Accept json
// @Param req body APIUpdateEventRequest true "Update parameters"
// @Param id path string true "Event id"
// @Success 200 {object} APIEvent
// @Router /crud/events/:id [put]
func (*GeneratedCrudController) UpdateEventByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	var updateReq UpdateEventRequest

	err := c.BindJSON(&updateReq)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	existingEvent, err := models.Events(qm.Where("id = ?", id), qm.For("UPDATE")).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	if updateReq.GameTypeID != nil {
		existingEvent.GameTypeID = *updateReq.GameTypeID
	}

	if updateReq.EventDate != nil {
		existingEvent.EventDate = *updateReq.EventDate
	}

	if updateReq.PlayerCount != nil {
		existingEvent.PlayerCount = *updateReq.PlayerCount
	}

	if updateReq.Name != nil {
		existingEvent.Name = *updateReq.Name
	}

	if updateReq.CreatedByID != nil {
		existingEvent.CreatedByID = *updateReq.CreatedByID
	}

	if updateReq.CourseID != nil {
		existingEvent.CourseID = *updateReq.CourseID
	}

	if updateReq.ParentID != nil {
		existingEvent.ParentID = *updateReq.ParentID
	}

	if updateReq.Public != nil {
		existingEvent.Public = *updateReq.Public
	}

	_, err = existingEvent.UpdateG(c.Request.Context(), boil.Infer())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, existingEvent)
}