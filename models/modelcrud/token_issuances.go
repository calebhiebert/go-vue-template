// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package modelcrud

import (
	"net/http"
	"time"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type APITokenIssuance struct {
	ID        string `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    string `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	IPAddress string `boil:"ip_address" json:"ip_address" toml:"ip_address" yaml:"ip_address"`

	CreatedAt *time.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
}

type GetTokenIssuancesResponse struct {
	TokenIssuances models.TokenIssuanceSlice `json:"token_issuances"`
	Total          int64                     `json:"total"`
	NextOffset     int64                     `json:"next_offset"`
}

type APIGetTokenIssuancesResponse struct {
	TokenIssuances []APITokenIssuance `json:"token_issuances"`
	Total          int64              `json:"total"`
	NextOffset     int64              `json:"next_offset"`
}

// GetTokenIssuanceByID godoc
// @Summary Gets a single TokenIssuance entity by their id
// @Produce json
// @Success 200 {object} APIGetTokenIssuancesResponse
// @Param id path string true "TokenIssuance id"
// @Router /crud/tokenIssuances/:id [get]
func (*GeneratedCrudController) GetTokenIssuanceByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	TokenIssuance, err := models.TokenIssuances(qm.Where("id = ?", id)).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, TokenIssuance)
}

// GetTokenIssuances godoc
// @Summary Gets a list for all entities of the TokenIssuance type
// @Produce json
// @Success 200 {object} APITokenIssuance
// @Router /crud/tokenIssuances [get]
func (*GeneratedCrudController) GetTokenIssuances(c *gin.Context) {
	limit, offset := api.ExtractLimitOffset(c)

	count, err := models.TokenIssuances().CountG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	queryMods := []qm.QueryMod{
		qm.Limit(limit),
		qm.Offset(offset),
	}

	tokenIssuances, err := models.TokenIssuances(queryMods...).AllG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	if tokenIssuances == nil {
		tokenIssuances = models.TokenIssuanceSlice{}
	}

	c.JSON(http.StatusOK, GetTokenIssuancesResponse{
		TokenIssuances: tokenIssuances,
		Total:          count,
		NextOffset:     int64(offset + limit),
	})
}

type APIUpdateTokenIssuanceRequest struct {
	UserID    *string `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	IPAddress *string `boil:"ip_address" json:"ip_address" toml:"ip_address" yaml:"ip_address"`
}

type UpdateTokenIssuanceRequest struct {
	UserID    *string `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id"`
	IPAddress *string `boil:"ip_address" json:"ip_address,omitempty" toml:"ip_address" yaml:"ip_address"`
}

// UpdateTokenIssuanceByID godoc
// @Summary Updates a single TokenIssuance entity based on their id
// @Produce json
// @Accept json
// @Param req body APIUpdateTokenIssuanceRequest true "Update parameters"
// @Param id path string true "TokenIssuance id"
// @Success 200 {object} APITokenIssuance
// @Router /crud/tokenIssuances/:id [put]
func (*GeneratedCrudController) UpdateTokenIssuanceByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	var updateReq UpdateTokenIssuanceRequest

	err := c.BindJSON(&updateReq)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	existingTokenIssuance, err := models.TokenIssuances(qm.Where("id = ?", id), qm.For("UPDATE")).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	if updateReq.UserID != nil {
		existingTokenIssuance.UserID = *updateReq.UserID
	}

	if updateReq.IPAddress != nil {
		existingTokenIssuance.IPAddress = *updateReq.IPAddress
	}

	_, err = existingTokenIssuance.UpdateG(c.Request.Context(), boil.Infer())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, existingTokenIssuance)
}

// DeleteTokenIssuanceByID godoc
// @Summary Soft deletes a single TokenIssuance entity based on their id
// @Produce json
// @Success 200 {object} APITokenIssuance
// @Param id path string true "TokenIssuance id"
// @Router /crud/tokenIssuances/:id [delete]
func (*GeneratedCrudController) DeleteTokenIssuanceByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	existingTokenIssuance, err := models.TokenIssuances(qm.Where("id = ?", id)).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	_, err = existingTokenIssuance.DeleteG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, existingTokenIssuance)
}

func (gcc *GeneratedCrudController) RegisterTokenIssuances(rg *gin.RouterGroup) {
	rg.GET("/tokenIssuances/:id", gcc.GetTokenIssuanceByID)
	rg.GET("/tokenIssuances", gcc.GetTokenIssuances)
	rg.PUT("/tokenIssuances/:id", gcc.UpdateTokenIssuanceByID)
	rg.DELETE("/tokenIssuances/:id", gcc.DeleteTokenIssuanceByID)
}

var TokenIssuancesAdmin = api.AdminModel{
	Name:          "TokenIssuances",
	CanSoftDelete: false,
	URLName:       "tokenIssuances",
	DataName:      "token_issuances",
	Fields: []*api.AdminModelField{&api.AdminModelField{
		ID:       "id",
		Name:     "ID",
		Nullable: false,
		Config:   api.NewDefaultAdminModelFieldConfig(),
		Type:     "string",
	}, &api.AdminModelField{
		ID:       "user_id",
		Name:     "UserID",
		Nullable: false,
		Config:   api.NewDefaultAdminModelFieldConfig(),
		Type:     "string",
	}, &api.AdminModelField{
		ID:       "ip_address",
		Name:     "IPAddress",
		Nullable: false,
		Config:   api.NewDefaultAdminModelFieldConfig(),
		Type:     "string",
	}, &api.AdminModelField{
		ID:       "created_at",
		Name:     "CreatedAt",
		Nullable: true,
		Config:   api.NewDefaultAdminModelFieldConfig(),
		Type:     "time",
	},
	},
}

type TokenIssuancesModelConfigType struct {
	ID        api.AdminModelFieldConfig
	UserID    api.AdminModelFieldConfig
	IPAddress api.AdminModelFieldConfig
	CreatedAt api.AdminModelFieldConfig
}

var TokenIssuancesModelConfig = TokenIssuancesModelConfigType{
	ID: api.AdminModelFieldConfig{
		ShowOnGraph: true,
	},
	UserID: api.AdminModelFieldConfig{
		ShowOnGraph: true,
	},
	IPAddress: api.AdminModelFieldConfig{
		ShowOnGraph: true,
	},
	CreatedAt: api.AdminModelFieldConfig{
		ShowOnGraph: true,
	},
}

func (c TokenIssuancesModelConfigType) Apply() {
	TokenIssuancesAdmin.Fields[0].Config = c.ID
	TokenIssuancesAdmin.Fields[1].Config = c.UserID
	TokenIssuancesAdmin.Fields[2].Config = c.IPAddress
	TokenIssuancesAdmin.Fields[3].Config = c.CreatedAt

}
