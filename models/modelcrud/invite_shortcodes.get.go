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

type APIInviteShortcode struct {
	// character varying
	ID string `boil:"id" json:"id" toml:"id" yaml:"id"`

	// uuid
	InviteID string `boil:"invite_id" json:"invite_id" toml:"invite_id" yaml:"invite_id"`

	// timestamp without time zone
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
}

type GetInviteShortcodesResponse struct {
	InviteShortcodes models.InviteShortcodeSlice `json:"invite_shortcodes"`
	Total            int64                       `json:"total"`
	NextOffset       int64                       `json:"next_offset"`
}

type APIGetInviteShortcodesResponse struct {
	InviteShortcodes []APIInviteShortcode `json:"invite_shortcodes"`
	Total            int64                `json:"total"`
	NextOffset       int64                `json:"next_offset"`
}

// GetInviteShortcodeByID godoc
// @Summary Gets a single InviteShortcode entity by their id
// @Produce json
// @Success 200 {object} APIGetInviteShortcodesResponse
// @Param id path string true "InviteShortcode id"
// @Router /crud/inviteShortcodes/:id [get]
func (*GeneratedCrudController) GetInviteShortcodeByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	InviteShortcode, err := models.InviteShortcodes(qm.Where("id = ?", id)).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, InviteShortcode)
}

// GetInviteShortcodes godoc
// @Summary Gets a list for all entities of the InviteShortcode type
// @Produce json
// @Success 200 {object} APIInviteShortcode
// @Param sort.id query string false "Sort by id. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.invite_id query string false "Sort by invite_id. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.created_at query string false "Sort by created_at. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Router /crud/inviteShortcodes [get]
func (*GeneratedCrudController) GetInviteShortcodes(c *gin.Context) {
	queryMods := []qm.QueryMod{}

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

		case "id.cont":
			idSearchString := fmt.Sprintf("%%%s%%", strings.ReplaceAll(v[0], "%", "\\%"))
			queryMods = append(queryMods, qm.Where("id ILIKE ?", idSearchString))

		case "sort.invite_id":
			orderBy = append(orderBy, "invite_id "+sortDirection)
		case "invite_id.eq":
			queryMods = append(queryMods, qm.Where("invite_id = ?", v[0]))
		case "invite_id.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("invite_id IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("invite_id IS NOT NULL"))
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

		}
	}

	count, err := models.InviteShortcodes(queryMods...).CountG(c.Request.Context())
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

	inviteShortcodes, err := models.InviteShortcodes(queryMods...).AllG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	if inviteShortcodes == nil {
		inviteShortcodes = models.InviteShortcodeSlice{}
	}

	c.JSON(http.StatusOK, GetInviteShortcodesResponse{
		InviteShortcodes: inviteShortcodes,
		Total:            count,
		NextOffset:       int64(offset + limit),
	})
}