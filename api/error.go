package api

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/gin-gonic/gin"
)

var MissingUserError = NewAPIError("missing-user", http.StatusUnauthorized, "User is not authenticated")
var MissingPermissionsError = NewAPIError("missing-permissions", http.StatusForbidden, "User is missing permissions")

type APIError struct {
	ID         string                 `json:"id"`
	Message    string                 `json:"message"`
	Code       int                    `json:"code"`
	Validation []ValidationFieldError `json:"validation,omitempty"`
	Details    interface{}            `json:"details"`
}

type ValidationFieldError struct {
	Tag     string      `json:"tag"`
	Field   string      `json:"field"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
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

	if e, ok := err.(*APIError); ok {
		return e
	}

	if e, ok := err.(validator.ValidationErrors); ok {
		var errs []ValidationFieldError

		for _, ve := range e {
			errs = append(errs, ValidationFieldError{
				Tag:     ve.Tag(),
				Field:   ve.Field(),
				Value:   ve.Value(),
				Message: ve.Error(),
			})
		}

		return &APIError{
			ID:         "validation-error",
			Code:       http.StatusBadRequest,
			Message:    "There was an error validating the input",
			Validation: errs,
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
	return fmt.Sprintf("(%s - %d): %s", e.ID, e.Code, e.Message)
}

func (e *APIError) Respond(c *gin.Context) {
	c.JSON(e.Code, e)
}

func (e *APIError) Abort(c *gin.Context) {
	c.AbortWithStatusJSON(e.Code, e)
}
