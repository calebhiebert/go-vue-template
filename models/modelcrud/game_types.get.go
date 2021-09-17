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

type APIGameType struct {
	// uuid
	ID string `boil:"id" json:"id" toml:"id" yaml:"id"`

	// character varying
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	// boolean
	IsCustom bool `boil:"is_custom" json:"is_custom" toml:"is_custom" yaml:"is_custom"`

	// timestamp without time zone

	DeletedAt *time.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
}

type GetGameTypesResponse struct {
	GameTypes  models.GameTypeSlice `json:"game_types"`
	Total      int64                `json:"total"`
	NextOffset int64                `json:"next_offset"`
}

type APIGetGameTypesResponse struct {
	GameTypes  []APIGameType `json:"game_types"`
	Total      int64         `json:"total"`
	NextOffset int64         `json:"next_offset"`
}

// GetGameTypeByID godoc
// @Summary Gets a single GameType entity by their id
// @Produce json
// @Success 200 {object} APIGetGameTypesResponse
// @Param id path string true "GameType id"
// @Router /crud/gameTypes/:id [get]
func (*GeneratedCrudController) GetGameTypeByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	GameType, err := models.GameTypes(qm.Where("id = ?", id)).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, GameType)
}

// GetGameTypes godoc
// @Summary Gets a list for all entities of the GameType type
// @Produce json
// @Success 200 {object} APIGameType
// @Param withDeleted query string false "Include deleted gameTypes in the results"
// @Param sort.id query string false "Sort by id. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.name query string false "Sort by name. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.is_custom query string false "Sort by is_custom. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.deleted_at query string false "Sort by deleted_at. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Router /crud/gameTypes [get]
func (*GeneratedCrudController) GetGameTypes(c *gin.Context) {
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

		case "sort.is_custom":
			orderBy = append(orderBy, "is_custom "+sortDirection)
		case "is_custom.eq":
			queryMods = append(queryMods, qm.Where("is_custom = ?", v[0]))
		case "is_custom.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("is_custom IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("is_custom IS NOT NULL"))
			}

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

	count, err := models.GameTypes(queryMods...).CountG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	limit, offset := api.ExtractLimitOffset(c)

	queryMods = append(queryMods, qm.Limit(limit), qm.Offset(offset))

	if len(orderBy) > 0 {
		queryMods = append(queryMods, qm.OrderBy(strings.Join(orderBy, ", ")))
	}

	gameTypes, err := models.GameTypes(queryMods...).AllG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	if gameTypes == nil {
		gameTypes = models.GameTypeSlice{}
	}

	c.JSON(http.StatusOK, GetGameTypesResponse{
		GameTypes:  gameTypes,
		Total:      count,
		NextOffset: int64(offset + limit),
	})
}