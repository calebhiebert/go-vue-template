// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package modelcrud

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type APIEvent struct {
	// uuid
	ID string `boil:"id" json:"id" toml:"id" yaml:"id"`

	// uuid
	GameTypeID string `boil:"game_type_id" json:"game_type_id" toml:"game_type_id" yaml:"game_type_id"`

	// timestamp without time zone
	EventDate time.Time `boil:"event_date" json:"event_date" toml:"event_date" yaml:"event_date"`

	// integer
	PlayerCount int `boil:"player_count" json:"player_count" toml:"player_count" yaml:"player_count"`

	// character varying
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	// uuid
	CreatedByID string `boil:"created_by_id" json:"created_by_id" toml:"created_by_id" yaml:"created_by_id"`

	// uuid
	CourseID string `boil:"course_id" json:"course_id" toml:"course_id" yaml:"course_id"`

	// uuid

	ParentID *string `boil:"parent_id" json:"parent_id,omitempty" toml:"parent_id" yaml:"parent_id,omitempty"`

	// boolean
	Public bool `boil:"public" json:"public" toml:"public" yaml:"public"`

	// timestamp without time zone
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	// timestamp without time zone
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	// timestamp without time zone

	DeletedAt *time.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
}

type GetEventsResponse struct {
	Events     models.EventSlice `json:"events"`
	Total      int64             `json:"total"`
	NextOffset int64             `json:"next_offset"`
}

type APIGetEventsResponse struct {
	Events     []APIEvent `json:"events"`
	Total      int64      `json:"total"`
	NextOffset int64      `json:"next_offset"`
}

// GetEventByID godoc
// @Summary Gets a single Event entity by their id
// @Produce json
// @Success 200 {object} APIGetEventsResponse
// @Param id path string true "Event id"
// @Router /crud/events/:id [get]
func (*GeneratedCrudController) GetEventByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	Event, err := models.Events(qm.Where("id = ?", id)).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, Event)
}

// GetEvents godoc
// @Summary Gets a list for all entities of the Event type
// @Produce json
// @Success 200 {object} APIEvent
// @Param withDeleted query string false "Include deleted events in the results"
// @Param sort.id query string false "Sort by id. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.game_type_id query string false "Sort by game_type_id. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.event_date query string false "Sort by event_date. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.player_count query string false "Sort by player_count. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.name query string false "Sort by name. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.created_by_id query string false "Sort by created_by_id. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.course_id query string false "Sort by course_id. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.parent_id query string false "Sort by parent_id. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.public query string false "Sort by public. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.created_at query string false "Sort by created_at. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.updated_at query string false "Sort by updated_at. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.deleted_at query string false "Sort by deleted_at. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Router /crud/events [get]
func (*GeneratedCrudController) GetEvents(c *gin.Context) {
	queryMods := []qm.QueryMod{}

	withDeleted := c.Query("withDeleted") == "true"

	if withDeleted {
		queryMods = append(queryMods, qm.WithDeleted())
	}

	var orderBy []string

	for q, v := range c.Request.URL.Query() {
		sortDirection := "ASC"

		if v[0] == "DESC" || v[0] == "desc" {
			sortDirection = "DESC"
		}

		switch q {
		case "sort.id":
			orderBy = append(orderBy, "id "+sortDirection)
		case "id.eq":
			queryMods = append(queryMods, qm.Where("id = ?", v[0]))
		case "id.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("id IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("id IS NOT NULL"))
			}

		case "sort.game_type_id":
			orderBy = append(orderBy, "game_type_id "+sortDirection)
		case "game_type_id.eq":
			queryMods = append(queryMods, qm.Where("game_type_id = ?", v[0]))
		case "game_type_id.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("game_type_id IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("game_type_id IS NOT NULL"))
			}

		case "sort.event_date":
			orderBy = append(orderBy, "event_date "+sortDirection)
		case "event_date.eq":
			queryMods = append(queryMods, qm.Where("event_date = ?", v[0]))
		case "event_date.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("event_date IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("event_date IS NOT NULL"))
			}

		case "event_date.gt":
			queryMods = append(queryMods, qm.Where("event_date > ?", v[0]))
		case "event_date.lt":
			queryMods = append(queryMods, qm.Where("event_date < ?", v[0]))
		case "event_date.gte":
			queryMods = append(queryMods, qm.Where("event_date >= ?", v[0]))
		case "event_date.lte":
			queryMods = append(queryMods, qm.Where("event_date <= ?", v[0]))

		case "sort.player_count":
			orderBy = append(orderBy, "player_count "+sortDirection)
		case "player_count.eq":
			queryMods = append(queryMods, qm.Where("player_count = ?", v[0]))
		case "player_count.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("player_count IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("player_count IS NOT NULL"))
			}

		case "player_count.gt":
			queryMods = append(queryMods, qm.Where("player_count > ?", v[0]))
		case "player_count.lt":
			queryMods = append(queryMods, qm.Where("player_count < ?", v[0]))
		case "player_count.gte":
			queryMods = append(queryMods, qm.Where("player_count >= ?", v[0]))
		case "player_count.lte":
			queryMods = append(queryMods, qm.Where("player_count <= ?", v[0]))

		case "sort.name":
			orderBy = append(orderBy, "name "+sortDirection)
		case "name.eq":
			queryMods = append(queryMods, qm.Where("name = ?", v[0]))
		case "name.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("name IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("name IS NOT NULL"))
			}

		case "name.cont":
			nameSearchString := fmt.Sprintf("%%%s%%", strings.ReplaceAll(v[0], "%", "\\%"))
			queryMods = append(queryMods, qm.Where("name ILIKE ?", nameSearchString))

		case "sort.created_by_id":
			orderBy = append(orderBy, "created_by_id "+sortDirection)
		case "created_by_id.eq":
			queryMods = append(queryMods, qm.Where("created_by_id = ?", v[0]))
		case "created_by_id.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("created_by_id IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("created_by_id IS NOT NULL"))
			}

		case "sort.course_id":
			orderBy = append(orderBy, "course_id "+sortDirection)
		case "course_id.eq":
			queryMods = append(queryMods, qm.Where("course_id = ?", v[0]))
		case "course_id.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("course_id IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("course_id IS NOT NULL"))
			}

		case "sort.parent_id":
			orderBy = append(orderBy, "parent_id "+sortDirection)
		case "parent_id.eq":
			queryMods = append(queryMods, qm.Where("parent_id = ?", v[0]))
		case "parent_id.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("parent_id IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("parent_id IS NOT NULL"))
			}

		case "sort.public":
			orderBy = append(orderBy, "public "+sortDirection)
		case "public.eq":
			queryMods = append(queryMods, qm.Where("public = ?", v[0]))
		case "public.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("public IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("public IS NOT NULL"))
			}

		case "sort.created_at":
			orderBy = append(orderBy, "created_at "+sortDirection)
		case "created_at.eq":
			queryMods = append(queryMods, qm.Where("created_at = ?", v[0]))
		case "created_at.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("created_at IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("created_at IS NOT NULL"))
			}

		case "created_at.gt":
			queryMods = append(queryMods, qm.Where("created_at > ?", v[0]))
		case "created_at.lt":
			queryMods = append(queryMods, qm.Where("created_at < ?", v[0]))
		case "created_at.gte":
			queryMods = append(queryMods, qm.Where("created_at >= ?", v[0]))
		case "created_at.lte":
			queryMods = append(queryMods, qm.Where("created_at <= ?", v[0]))

		case "sort.updated_at":
			orderBy = append(orderBy, "updated_at "+sortDirection)
		case "updated_at.eq":
			queryMods = append(queryMods, qm.Where("updated_at = ?", v[0]))
		case "updated_at.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("updated_at IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("updated_at IS NOT NULL"))
			}

		case "updated_at.gt":
			queryMods = append(queryMods, qm.Where("updated_at > ?", v[0]))
		case "updated_at.lt":
			queryMods = append(queryMods, qm.Where("updated_at < ?", v[0]))
		case "updated_at.gte":
			queryMods = append(queryMods, qm.Where("updated_at >= ?", v[0]))
		case "updated_at.lte":
			queryMods = append(queryMods, qm.Where("updated_at <= ?", v[0]))

		case "sort.deleted_at":
			orderBy = append(orderBy, "deleted_at "+sortDirection)
		case "deleted_at.eq":
			queryMods = append(queryMods, qm.Where("deleted_at = ?", v[0]))
		case "deleted_at.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("deleted_at IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("deleted_at IS NOT NULL"))
			}

		case "deleted_at.gt":
			queryMods = append(queryMods, qm.Where("deleted_at > ?", v[0]))
		case "deleted_at.lt":
			queryMods = append(queryMods, qm.Where("deleted_at < ?", v[0]))
		case "deleted_at.gte":
			queryMods = append(queryMods, qm.Where("deleted_at >= ?", v[0]))
		case "deleted_at.lte":
			queryMods = append(queryMods, qm.Where("deleted_at <= ?", v[0]))

		}
	}

	count, err := models.Events(queryMods...).CountG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	limit, offset := api.ExtractLimitOffset(c)

	queryMods = append(queryMods, qm.Limit(limit), qm.Offset(offset))

	if len(orderBy) > 0 {
		queryMods = append(queryMods, qm.OrderBy(strings.Join(orderBy, ", ")))
	} else {
		queryMods = append(queryMods, qm.OrderBy("created_at DESC"))
	}

	events, err := models.Events(queryMods...).AllG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	if events == nil {
		events = models.EventSlice{}
	}

	c.JSON(http.StatusOK, GetEventsResponse{
		Events:     events,
		Total:      count,
		NextOffset: int64(offset + limit),
	})
}
