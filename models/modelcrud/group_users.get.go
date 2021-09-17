// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package modelcrud

import (
	"net/http"
	"strings"
	"time"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type APIGroupUser struct {
	// uuid
	GroupID string `boil:"group_id" json:"group_id" toml:"group_id" yaml:"group_id"`

	// uuid
	UserID string `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`

	// timestamp without time zone
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	// timestamp without time zone
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	// timestamp without time zone

	DeletedAt *time.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
}

type GetGroupUsersResponse struct {
	GroupUsers models.GroupUserSlice `json:"group_users"`
	Total      int64                 `json:"total"`
	NextOffset int64                 `json:"next_offset"`
}

type APIGetGroupUsersResponse struct {
	GroupUsers []APIGroupUser `json:"group_users"`
	Total      int64          `json:"total"`
	NextOffset int64          `json:"next_offset"`
}

// GetGroupUserByID godoc
// @Summary Gets a single GroupUser entity by their id
// @Produce json
// @Success 200 {object} APIGetGroupUsersResponse
// @Param id path string true "GroupUser id"
// @Router /crud/groupUsers/:id [get]
func (*GeneratedCrudController) GetGroupUserByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	GroupUser, err := models.GroupUsers(qm.Where("id = ?", id)).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, GroupUser)
}

// GetGroupUsers godoc
// @Summary Gets a list for all entities of the GroupUser type
// @Produce json
// @Success 200 {object} APIGroupUser
// @Param withDeleted query string false "Include deleted groupUsers in the results"
// @Param sort.group_id query string false "Sort by group_id. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.user_id query string false "Sort by user_id. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.created_at query string false "Sort by created_at. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.updated_at query string false "Sort by updated_at. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.deleted_at query string false "Sort by deleted_at. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Router /crud/groupUsers [get]
func (*GeneratedCrudController) GetGroupUsers(c *gin.Context) {
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
		case "sort.group_id":
			orderBy = append(orderBy, "group_id "+sortDirection)
		case "group_id.eq":
			queryMods = append(queryMods, qm.Where("group_id = ?", v[0]))
		case "group_id.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("group_id IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("group_id IS NOT NULL"))
			}

		case "sort.user_id":
			orderBy = append(orderBy, "user_id "+sortDirection)
		case "user_id.eq":
			queryMods = append(queryMods, qm.Where("user_id = ?", v[0]))
		case "user_id.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("user_id IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("user_id IS NOT NULL"))
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

	count, err := models.GroupUsers(queryMods...).CountG(c.Request.Context())
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

	groupUsers, err := models.GroupUsers(queryMods...).AllG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	if groupUsers == nil {
		groupUsers = models.GroupUserSlice{}
	}

	c.JSON(http.StatusOK, GetGroupUsersResponse{
		GroupUsers: groupUsers,
		Total:      count,
		NextOffset: int64(offset + limit),
	})
}