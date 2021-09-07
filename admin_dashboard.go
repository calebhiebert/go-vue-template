package main

import (
	"net/http"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetAdminDashboardStats(c *gin.Context) {
	userCount, err := models.Users().CountG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	newUsersMonthly, err := models.Users(qm.
		Where("created_at BETWEEN (now() at time zone 'utc' - interval '1 month') AND (now() at time zone 'utc')")).
		CountG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	newUsersWeekly, err := models.Users(qm.
		Where("created_at BETWEEN (now() at time zone 'utc' - interval '1 week') AND (now() at time zone 'utc')")).
		CountG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_users": userCount,
		"monthly_new_users": newUsersMonthly,
		"weekly_new_users": newUsersWeekly,
	})
}
