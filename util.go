package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func extractLimitOffset(c *gin.Context) (limit int, offset int) {
	limit, _ = strconv.Atoi(c.Query("limit"))
	offset, _ = strconv.Atoi(c.Query("offset"))

	if limit < 0 {
		limit = 0
	} else if limit > 100 {
		limit = 100
	}

	if offset < 0 {
		offset = 0
	}

	return
}

func stringSliceContains(s []string, c ...string) bool {
	for _, str := range s {
		for _, chk := range c {
			if chk == str {
				return true
			}
		}
	}

	return false
}