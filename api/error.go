package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	ID      string      `json:"id"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Details interface{} `json:"details"`
}

func NewAPIError(id string, code int, message string) *APIError {
	return &APIError{
		ID:      id,
		Code:    code,
		Message: message,
	}
}

func APIErrorFromErr(err error) *APIError {
	if err == sql.ErrNoRows {
		return &APIError{
			ID:      "not-found",
			Code:    http.StatusNotFound,
			Message: "The requested resource could not be found",
		}
	}

	return &APIError{
		ID:      "unknown_error",
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
		Details: err,
	}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("(%s - %s): %s", e.ID, e.Code, e.Message)
}

func (e *APIError) Respond(c *gin.Context) {
	c.JSON(e.Code, e)
}

func (e *APIError) Abort(c *gin.Context) {
	c.AbortWithStatusJSON(e.Code, e)
}
