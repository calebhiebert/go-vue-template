// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package modelcrud

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type APIUserProfileQuestion struct {
	// uuid
	UserID string `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`

	// character varying
	ProfileQuestionID string `boil:"profile_question_id" json:"profile_question_id" toml:"profile_question_id" yaml:"profile_question_id"`

	// integer

	ANumber *int `boil:"a_number" json:"a_number,omitempty" toml:"a_number" yaml:"a_number,omitempty"`

	// integer

	ARating *int `boil:"a_rating" json:"a_rating,omitempty" toml:"a_rating" yaml:"a_rating,omitempty"`

	// character varying

	ADays *string `boil:"a_days" json:"a_days,omitempty" toml:"a_days" yaml:"a_days,omitempty"`

	// text

	AText *string `boil:"a_text" json:"a_text,omitempty" toml:"a_text" yaml:"a_text,omitempty"`
}

type GetUserProfileQuestionsResponse struct {
	UserProfileQuestions models.UserProfileQuestionSlice `json:"user_profile_questions"`
	Total                int64                           `json:"total"`
	NextOffset           int64                           `json:"next_offset"`
}

type APIGetUserProfileQuestionsResponse struct {
	UserProfileQuestions []APIUserProfileQuestion `json:"user_profile_questions"`
	Total                int64                    `json:"total"`
	NextOffset           int64                    `json:"next_offset"`
}

// GetUserProfileQuestionByID godoc
// @Summary Gets a single UserProfileQuestion entity by their id
// @Produce json
// @Success 200 {object} APIGetUserProfileQuestionsResponse
// @Param id path string true "UserProfileQuestion id"
// @Router /crud/userProfileQuestions/:id [get]
func (*GeneratedCrudController) GetUserProfileQuestionByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	UserProfileQuestion, err := models.UserProfileQuestions(qm.Where("id = ?", id)).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, UserProfileQuestion)
}

// GetUserProfileQuestions godoc
// @Summary Gets a list for all entities of the UserProfileQuestion type
// @Produce json
// @Success 200 {object} APIUserProfileQuestion
// @Param sort.user_id query string false "Sort by user_id. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.profile_question_id query string false "Sort by profile_question_id. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.a_number query string false "Sort by a_number. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.a_rating query string false "Sort by a_rating. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.a_days query string false "Sort by a_days. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.a_text query string false "Sort by a_text. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Router /crud/userProfileQuestions [get]
func (*GeneratedCrudController) GetUserProfileQuestions(c *gin.Context) {
	queryMods := []qm.QueryMod{}

	var orderBy []string

	for q, v := range c.Request.URL.Query() {
		sortDirection := "ASC"

		if v[0] == "DESC" || v[0] == "desc" {
			sortDirection = "DESC"
		}

		switch q {
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

		case "sort.profile_question_id":
			orderBy = append(orderBy, "profile_question_id "+sortDirection)
		case "profile_question_id.eq":
			queryMods = append(queryMods, qm.Where("profile_question_id = ?", v[0]))
		case "profile_question_id.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("profile_question_id IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("profile_question_id IS NOT NULL"))
			}

		case "profile_question_id.cont":
			profile_question_idSearchString := fmt.Sprintf("%%%s%%", strings.ReplaceAll(v[0], "%", "\\%"))
			queryMods = append(queryMods, qm.Where("profile_question_id ILIKE ?", profile_question_idSearchString))

		case "sort.a_number":
			orderBy = append(orderBy, "a_number "+sortDirection)
		case "a_number.eq":
			queryMods = append(queryMods, qm.Where("a_number = ?", v[0]))
		case "a_number.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("a_number IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("a_number IS NOT NULL"))
			}

		case "a_number.gt":
			queryMods = append(queryMods, qm.Where("a_number > ?", v[0]))
		case "a_number.lt":
			queryMods = append(queryMods, qm.Where("a_number < ?", v[0]))
		case "a_number.gte":
			queryMods = append(queryMods, qm.Where("a_number >= ?", v[0]))
		case "a_number.lte":
			queryMods = append(queryMods, qm.Where("a_number <= ?", v[0]))

		case "sort.a_rating":
			orderBy = append(orderBy, "a_rating "+sortDirection)
		case "a_rating.eq":
			queryMods = append(queryMods, qm.Where("a_rating = ?", v[0]))
		case "a_rating.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("a_rating IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("a_rating IS NOT NULL"))
			}

		case "a_rating.gt":
			queryMods = append(queryMods, qm.Where("a_rating > ?", v[0]))
		case "a_rating.lt":
			queryMods = append(queryMods, qm.Where("a_rating < ?", v[0]))
		case "a_rating.gte":
			queryMods = append(queryMods, qm.Where("a_rating >= ?", v[0]))
		case "a_rating.lte":
			queryMods = append(queryMods, qm.Where("a_rating <= ?", v[0]))

		case "sort.a_days":
			orderBy = append(orderBy, "a_days "+sortDirection)
		case "a_days.eq":
			queryMods = append(queryMods, qm.Where("a_days = ?", v[0]))
		case "a_days.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("a_days IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("a_days IS NOT NULL"))
			}

		case "a_days.cont":
			a_daysSearchString := fmt.Sprintf("%%%s%%", strings.ReplaceAll(v[0], "%", "\\%"))
			queryMods = append(queryMods, qm.Where("a_days ILIKE ?", a_daysSearchString))

		case "sort.a_text":
			orderBy = append(orderBy, "a_text "+sortDirection)
		case "a_text.eq":
			queryMods = append(queryMods, qm.Where("a_text = ?", v[0]))
		case "a_text.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("a_text IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("a_text IS NOT NULL"))
			}

		case "a_text.cont":
			a_textSearchString := fmt.Sprintf("%%%s%%", strings.ReplaceAll(v[0], "%", "\\%"))
			queryMods = append(queryMods, qm.Where("a_text ILIKE ?", a_textSearchString))
		}
	}

	count, err := models.UserProfileQuestions(queryMods...).CountG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	limit, offset := api.ExtractLimitOffset(c)

	queryMods = append(queryMods, qm.Limit(limit), qm.Offset(offset))

	if len(orderBy) > 0 {
		queryMods = append(queryMods, qm.OrderBy(strings.Join(orderBy, ", ")))
	}

	userProfileQuestions, err := models.UserProfileQuestions(queryMods...).AllG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	if userProfileQuestions == nil {
		userProfileQuestions = models.UserProfileQuestionSlice{}
	}

	c.JSON(http.StatusOK, GetUserProfileQuestionsResponse{
		UserProfileQuestions: userProfileQuestions,
		Total:                count,
		NextOffset:           int64(offset + limit),
	})
}