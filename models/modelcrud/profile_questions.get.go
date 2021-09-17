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

type APIProfileQuestion struct {
	// character varying
	ID string `boil:"id" json:"id" toml:"id" yaml:"id"`

	// text
	Question string `boil:"question" json:"question" toml:"question" yaml:"question"`

	// integer
	AnswerType int `boil:"answer_type" json:"answer_type" toml:"answer_type" yaml:"answer_type"`

	// jsonb

	Options map[string]interface{} `boil:"options" json:"options" toml:"options" yaml:"options"`

	// double precision
	Weight float64 `boil:"weight" json:"weight" toml:"weight" yaml:"weight"`

	// timestamp without time zone

	CreatedAt *time.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
}

type GetProfileQuestionsResponse struct {
	ProfileQuestions models.ProfileQuestionSlice `json:"profile_questions"`
	Total            int64                       `json:"total"`
	NextOffset       int64                       `json:"next_offset"`
}

type APIGetProfileQuestionsResponse struct {
	ProfileQuestions []APIProfileQuestion `json:"profile_questions"`
	Total            int64                `json:"total"`
	NextOffset       int64                `json:"next_offset"`
}

// GetProfileQuestionByID godoc
// @Summary Gets a single ProfileQuestion entity by their id
// @Produce json
// @Success 200 {object} APIGetProfileQuestionsResponse
// @Param id path string true "ProfileQuestion id"
// @Router /crud/profileQuestions/:id [get]
func (*GeneratedCrudController) GetProfileQuestionByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	ProfileQuestion, err := models.ProfileQuestions(qm.Where("id = ?", id)).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, ProfileQuestion)
}

// GetProfileQuestions godoc
// @Summary Gets a list for all entities of the ProfileQuestion type
// @Produce json
// @Success 200 {object} APIProfileQuestion
// @Param sort.id query string false "Sort by id. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.question query string false "Sort by question. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.answer_type query string false "Sort by answer_type. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.options query string false "Sort by options. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.weight query string false "Sort by weight. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.created_at query string false "Sort by created_at. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Router /crud/profileQuestions [get]
func (*GeneratedCrudController) GetProfileQuestions(c *gin.Context) {
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

		case "sort.question":
			orderBy = append(orderBy, "question "+sortDirection)
		case "question.eq":
			queryMods = append(queryMods, qm.Where("question = ?", v[0]))
		case "question.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("question IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("question IS NOT NULL"))
			}

		case "question.cont":
			questionSearchString := fmt.Sprintf("%%%s%%", strings.ReplaceAll(v[0], "%", "\\%"))
			queryMods = append(queryMods, qm.Where("question ILIKE ?", questionSearchString))

		case "sort.answer_type":
			orderBy = append(orderBy, "answer_type "+sortDirection)
		case "answer_type.eq":
			queryMods = append(queryMods, qm.Where("answer_type = ?", v[0]))
		case "answer_type.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("answer_type IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("answer_type IS NOT NULL"))
			}

		case "answer_type.gt":
			queryMods = append(queryMods, qm.Where("answer_type > ?", v[0]))
		case "answer_type.lt":
			queryMods = append(queryMods, qm.Where("answer_type < ?", v[0]))
		case "answer_type.gte":
			queryMods = append(queryMods, qm.Where("answer_type >= ?", v[0]))
		case "answer_type.lte":
			queryMods = append(queryMods, qm.Where("answer_type <= ?", v[0]))

		case "sort.options":
			orderBy = append(orderBy, "options "+sortDirection)
		case "options.eq":
			queryMods = append(queryMods, qm.Where("options = ?", v[0]))
		case "options.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("options IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("options IS NOT NULL"))
			}

		case "sort.weight":
			orderBy = append(orderBy, "weight "+sortDirection)
		case "weight.eq":
			queryMods = append(queryMods, qm.Where("weight = ?", v[0]))
		case "weight.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("weight IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("weight IS NOT NULL"))
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

	count, err := models.ProfileQuestions(queryMods...).CountG(c.Request.Context())
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

	profileQuestions, err := models.ProfileQuestions(queryMods...).AllG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	if profileQuestions == nil {
		profileQuestions = models.ProfileQuestionSlice{}
	}

	c.JSON(http.StatusOK, GetProfileQuestionsResponse{
		ProfileQuestions: profileQuestions,
		Total:            count,
		NextOffset:       int64(offset + limit),
	})
}