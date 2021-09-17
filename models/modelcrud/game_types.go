// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package modelcrud

import (
	"github.com/calebhiebert/go-vue-template/api"
	"github.com/gin-gonic/gin"
)

func (gcc *GeneratedCrudController) RegisterGameTypes(rg *gin.RouterGroup) {
	rg.GET("/gameTypes/:id", gcc.GetGameTypeByID)
	rg.GET("/gameTypes", gcc.GetGameTypes)
	rg.PUT("/gameTypes/:id", gcc.UpdateGameTypeByID)
	rg.POST("/gameTypes", gcc.CreateGameType)
	rg.DELETE("/gameTypes/:id", gcc.DeleteGameTypeByID)
	rg.DELETE("/gameTypes", gcc.BulkDeleteGameTypesByIDs)
	rg.POST("/gameTypes/:id/unDelete", gcc.UnDeleteGameTypeByID)
}

var GameTypesAdmin = api.AdminModel{
	Name:          "GameTypes",
	NameSingular:  "GameType",
	CanSoftDelete: true,
	URLName:       "gameTypes",
	DataName:      "game_types",
	Fields: []*api.AdminModelField{
		&api.AdminModelField{
			ID:       "id",
			Name:     "ID",
			Nullable: false,
			Required: true,
			FilterOperations: []string{
				"eq", "null"},
			ForeignFields: []api.AdminModelForeignField{},
			Editable:      false,
			Config: api.AdminModelFieldConfig{
				ShowOnGraph: true,
				Editable:    true,
				IsEmail:     false,
			},
			Type: "uuid",
		},
		&api.AdminModelField{
			ID:       "name",
			Name:     "Name",
			Nullable: false,
			Required: true,
			FilterOperations: []string{
				"eq", "null", "cont"},
			ForeignFields: []api.AdminModelForeignField{},
			Editable:      true,
			Config: api.AdminModelFieldConfig{
				ShowOnGraph: true,
				Editable:    true,
				IsEmail:     false,
			},
			Type: "string",
		},
		&api.AdminModelField{
			ID:       "is_custom",
			Name:     "IsCustom",
			Nullable: false,
			Required: false,
			FilterOperations: []string{
				"eq", "null"},
			ForeignFields: []api.AdminModelForeignField{},
			Editable:      true,
			Config: api.AdminModelFieldConfig{
				ShowOnGraph: true,
				Editable:    true,
				IsEmail:     false,
			},
			Type: "bool",
		},
		&api.AdminModelField{
			ID:       "deleted_at",
			Name:     "DeletedAt",
			Nullable: true,
			Required: false,
			FilterOperations: []string{
				"eq", "null", "gt", "lt", "gte", "lte"},
			ForeignFields: []api.AdminModelForeignField{},
			Editable:      false,
			Config: api.AdminModelFieldConfig{
				ShowOnGraph: false,
				Editable:    true,
				IsEmail:     false,
			},
			Type: "time",
		},
	},
}

type GameTypesModelConfigType struct {
	ID api.AdminModelFieldConfig

	Name api.AdminModelFieldConfig

	IsCustom api.AdminModelFieldConfig

	DeletedAt api.AdminModelFieldConfig
}

var GameTypesModelConfig = GameTypesModelConfigType{

	ID: api.AdminModelFieldConfig{
		ShowOnGraph: true,
		Editable:    true,
		IsEmail:     false,
	},
	Name: api.AdminModelFieldConfig{
		ShowOnGraph: true,
		Editable:    true,
		IsEmail:     false,
	},
	IsCustom: api.AdminModelFieldConfig{
		ShowOnGraph: true,
		Editable:    true,
		IsEmail:     false,
	},
	DeletedAt: api.AdminModelFieldConfig{
		ShowOnGraph: false,
		Editable:    true,
		IsEmail:     false,
	},
}

func (c GameTypesModelConfigType) Apply() {

	GameTypesAdmin.Fields[0].Config = c.ID

	GameTypesAdmin.Fields[1].Config = c.Name

	GameTypesAdmin.Fields[2].Config = c.IsCustom

	GameTypesAdmin.Fields[3].Config = c.DeletedAt

}