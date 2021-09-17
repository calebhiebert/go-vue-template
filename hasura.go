package main

import (
	"fmt"
	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/util"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func HasuraProxy(c *gin.Context) {
	user := extractVerifiedUser(c)

	var role string

	if user == nil {
		role = "anon"
	} else {
		if util.StringSliceContains(user.Roles, "admin") {
			role = "admin"
		} else {
			role = "user"
		}
	}

	defer c.Request.Body.Close()

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(c.Request.Context(), "POST", os.Getenv("HASURA_URL"), c.Request.Body)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	req.Header.Add("X-Hasura-Role", role)
	req.Header.Add("X-Hasura-Allowed-Roles", "{admin,user,anon}")
	req.Header.Add("X-Hasura-Authorization", c.GetHeader("Authorization"))

	if user != nil {
		req.Header.Add("X-Hasura-User-ID", user.ID)
	}

	for h, v := range c.Request.Header {
		for _, value := range v {
			switch strings.ToLower(h) {
			case "connection":
				// Do nothing
			default:
				req.Header.Add(h, value)
			}
		}
	}

	res, err := client.Do(req)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	defer res.Body.Close()

	for h, v := range res.Header {
		if strings.ToLower(h) != strings.ToLower("Transfer-Encoding") {
			for _, value := range v {
				c.Writer.Header().Add(h, value)
			}
		}
	}

	c.Writer.WriteHeader(res.StatusCode)

	_, err = io.Copy(c.Writer, res.Body)
	if err != nil {
		fmt.Println("FAILED TO WRITE BODY", err.Error())
	}
}
