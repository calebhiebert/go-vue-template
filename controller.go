package main

import (
	"net/http"

	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
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
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Test godoc
// @Summary Run a test handler
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /test [get]
func (*Controller) Test(c *gin.Context) {
	test, err := models.Users().AllG(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	test2, err := models.TestTables().OneG(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"test": test, "test2": test2})
}


