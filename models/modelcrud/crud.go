// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package modelcrud

import "github.com/gin-gonic/gin"

type GeneratedCrudController struct{}

type IDList struct {
	IDs []string `json:"ids"`
}

type DeletedCount struct {
	DeletedCount int `json:"deleted_count"`
}

func RegisterGeneratedCrud(rg *gin.RouterGroup) {
	gcc := GeneratedCrudController{}

	gcc.RegisterAccessLogs(rg)
	gcc.RegisterImages(rg)
	gcc.RegisterJobs(rg)
	gcc.RegisterTokenIssuances(rg)
	gcc.RegisterUsers(rg)

	rg.GET("/_schema", func(c *gin.Context) {
		c.JSON(200, AdminInfo)
	})
}
