package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Controller struct {

}

func NewController() *Controller {
	return &Controller{}
}

// HealthCheck godoc
// @Summary Health check handler
// @Description returns 200 when service is running
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /healthz [get]
func (*Controller) HealthCheck(c *gin.Context) {
	err := boil.GetDB().(*sql.DB).PingContext(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

