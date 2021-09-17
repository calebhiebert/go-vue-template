// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package modelcrud

import (
	"net/http"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type GenerateEventPQMatchScoreStringID func(*CreateEventPQMatchScoreRequest) string

var GenerateEventPQMatchScoreID GenerateEventPQMatchScoreStringID

type APICreateEventPQMatchScoreRequest struct {
	EventID           *string `boil:"event_id" json:"event_id" toml:"event_id" yaml:"event_id"`
	ProfileQuestionID *string `boil:"profile_question_id" json:"profile_question_id" toml:"profile_question_id" yaml:"profile_question_id"`

	RatingAvg *int `boil:"rating_avg" json:"rating_avg,omitempty" toml:"rating_avg" yaml:"rating_avg,omitempty"`

	NumAvg *float64 `boil:"num_avg" json:"num_avg,omitempty" toml:"num_avg" yaml:"num_avg,omitempty"`
}

type CreateEventPQMatchScoreRequest struct {
	EventID           string        `boil:"event_id" json:"event_id,omitempty" toml:"event_id" binding:"required" yaml:"event_id"`
	ProfileQuestionID string        `boil:"profile_question_id" json:"profile_question_id,omitempty" toml:"profile_question_id" binding:"required" yaml:"profile_question_id"`
	RatingAvg         *null.Int     `boil:"rating_avg" json:"rating_avg,omitempty" toml:"rating_avg"  yaml:"rating_avg,omitempty"`
	NumAvg            *null.Float64 `boil:"num_avg" json:"num_avg,omitempty" toml:"num_avg"  yaml:"num_avg,omitempty"`
}

// CreateEventPQMatchScore godoc
// @Summary Creates a new EventPQMatchScore
// @Produce json
// @Accept json
// @Param req body APICreateEventPQMatchScoreRequest true "Creation parameters"
// @Success 200 {object} APIEventPQMatchScore
// @Router /crud/eventPQMatchScores  [post]
func (*GeneratedCrudController) CreateEventPQMatchScore(c *gin.Context) {
	var createReq CreateEventPQMatchScoreRequest

	err := c.BindJSON(&createReq)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	newEventPQMatchScore := models.EventPQMatchScore{}
	newEventPQMatchScore.EventID = createReq.EventID

	newEventPQMatchScore.ProfileQuestionID = createReq.ProfileQuestionID

	if createReq.RatingAvg != nil {
		newEventPQMatchScore.RatingAvg = *createReq.RatingAvg
	}

	if createReq.NumAvg != nil {
		newEventPQMatchScore.NumAvg = *createReq.NumAvg
	}

	err = newEventPQMatchScore.InsertG(c.Request.Context(), boil.Infer())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, newEventPQMatchScore)
}