// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package modelcrud

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type APICourse struct {
	// uuid
	ID string `boil:"id" json:"id" toml:"id" yaml:"id"`

	// character varying
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	// character varying

	Address *string `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`

	// double precision

	Latitude *float64 `boil:"latitude" json:"latitude,omitempty" toml:"latitude" yaml:"latitude,omitempty"`

	// double precision

	Longitude *float64 `boil:"longitude" json:"longitude,omitempty" toml:"longitude" yaml:"longitude,omitempty"`

	// timestamp without time zone
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	// timestamp without time zone
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	// timestamp without time zone

	DeletedAt *time.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
}

type GetCoursesResponse struct {
	Courses    models.CourseSlice `json:"courses"`
	Total      int64              `json:"total"`
	NextOffset int64              `json:"next_offset"`
}

type APIGetCoursesResponse struct {
	Courses    []APICourse `json:"courses"`
	Total      int64       `json:"total"`
	NextOffset int64       `json:"next_offset"`
}

// GetCourseByID godoc
// @Summary Gets a single Course entity by their id
// @Produce json
// @Success 200 {object} APIGetCoursesResponse
// @Param id path string true "Course id"
// @Router /crud/courses/:id [get]
func (*GeneratedCrudController) GetCourseByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		api.NewAPIError("invalid-id", http.StatusBadRequest, "The provided id was invalid").Respond(c)
		return
	}

	Course, err := models.Courses(qm.Where("id = ?", id)).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, Course)
}

// GetCourses godoc
// @Summary Gets a list for all entities of the Course type
// @Produce json
// @Success 200 {object} APICourse
// @Param withDeleted query string false "Include deleted courses in the results"
// @Param sort.id query string false "Sort by id. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.name query string false "Sort by name. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.address query string false "Sort by address. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.latitude query string false "Sort by latitude. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.longitude query string false "Sort by longitude. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.created_at query string false "Sort by created_at. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.updated_at query string false "Sort by updated_at. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Param sort.deleted_at query string false "Sort by deleted_at. Value should be ASC or DESC. eg: ?sort.created_at=DESC"
// @Router /crud/courses [get]
func (*GeneratedCrudController) GetCourses(c *gin.Context) {
	queryMods := []qm.QueryMod{}

	withDeleted := c.Query("withDeleted") == "true"

	if withDeleted {
		queryMods = append(queryMods, qm.WithDeleted())
	}

	var orderBy []string

	for q, v := range c.Request.URL.Query() {
		sortDirection := "ASC"

		if v[0] == "DESC" || v[0] == "desc" {
			sortDirection = "DESC"
		}

		switch q {
		case "sort.id":
			orderBy = append(orderBy, "id "+sortDirection)
		case "id.eq":
			queryMods = append(queryMods, qm.Where("id = ?", v[0]))
		case "id.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("id IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("id IS NOT NULL"))
			}

		case "sort.name":
			orderBy = append(orderBy, "name "+sortDirection)
		case "name.eq":
			queryMods = append(queryMods, qm.Where("name = ?", v[0]))
		case "name.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("name IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("name IS NOT NULL"))
			}

		case "name.cont":
			nameSearchString := fmt.Sprintf("%%%s%%", strings.ReplaceAll(v[0], "%", "\\%"))
			queryMods = append(queryMods, qm.Where("name ILIKE ?", nameSearchString))

		case "sort.address":
			orderBy = append(orderBy, "address "+sortDirection)
		case "address.eq":
			queryMods = append(queryMods, qm.Where("address = ?", v[0]))
		case "address.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("address IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("address IS NOT NULL"))
			}

		case "address.cont":
			addressSearchString := fmt.Sprintf("%%%s%%", strings.ReplaceAll(v[0], "%", "\\%"))
			queryMods = append(queryMods, qm.Where("address ILIKE ?", addressSearchString))

		case "sort.latitude":
			orderBy = append(orderBy, "latitude "+sortDirection)
		case "latitude.eq":
			queryMods = append(queryMods, qm.Where("latitude = ?", v[0]))
		case "latitude.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("latitude IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("latitude IS NOT NULL"))
			}

		case "latitude.gt":
			queryMods = append(queryMods, qm.Where("latitude > ?", v[0]))
		case "latitude.lt":
			queryMods = append(queryMods, qm.Where("latitude < ?", v[0]))
		case "latitude.gte":
			queryMods = append(queryMods, qm.Where("latitude >= ?", v[0]))
		case "latitude.lte":
			queryMods = append(queryMods, qm.Where("latitude <= ?", v[0]))

		case "sort.longitude":
			orderBy = append(orderBy, "longitude "+sortDirection)
		case "longitude.eq":
			queryMods = append(queryMods, qm.Where("longitude = ?", v[0]))
		case "longitude.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("longitude IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("longitude IS NOT NULL"))
			}

		case "longitude.gt":
			queryMods = append(queryMods, qm.Where("longitude > ?", v[0]))
		case "longitude.lt":
			queryMods = append(queryMods, qm.Where("longitude < ?", v[0]))
		case "longitude.gte":
			queryMods = append(queryMods, qm.Where("longitude >= ?", v[0]))
		case "longitude.lte":
			queryMods = append(queryMods, qm.Where("longitude <= ?", v[0]))

		case "sort.created_at":
			orderBy = append(orderBy, "created_at "+sortDirection)
		case "created_at.eq":
			queryMods = append(queryMods, qm.Where("created_at = ?", v[0]))
		case "created_at.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("created_at IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("created_at IS NOT NULL"))
			}

		case "created_at.gt":
			queryMods = append(queryMods, qm.Where("created_at > ?", v[0]))
		case "created_at.lt":
			queryMods = append(queryMods, qm.Where("created_at < ?", v[0]))
		case "created_at.gte":
			queryMods = append(queryMods, qm.Where("created_at >= ?", v[0]))
		case "created_at.lte":
			queryMods = append(queryMods, qm.Where("created_at <= ?", v[0]))

		case "sort.updated_at":
			orderBy = append(orderBy, "updated_at "+sortDirection)
		case "updated_at.eq":
			queryMods = append(queryMods, qm.Where("updated_at = ?", v[0]))
		case "updated_at.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("updated_at IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("updated_at IS NOT NULL"))
			}

		case "updated_at.gt":
			queryMods = append(queryMods, qm.Where("updated_at > ?", v[0]))
		case "updated_at.lt":
			queryMods = append(queryMods, qm.Where("updated_at < ?", v[0]))
		case "updated_at.gte":
			queryMods = append(queryMods, qm.Where("updated_at >= ?", v[0]))
		case "updated_at.lte":
			queryMods = append(queryMods, qm.Where("updated_at <= ?", v[0]))

		case "sort.deleted_at":
			orderBy = append(orderBy, "deleted_at "+sortDirection)
		case "deleted_at.eq":
			queryMods = append(queryMods, qm.Where("deleted_at = ?", v[0]))
		case "deleted_at.null":
			if v[0] == "true" {
				queryMods = append(queryMods, qm.Where("deleted_at IS NULL"))
			} else {
				queryMods = append(queryMods, qm.Where("deleted_at IS NOT NULL"))
			}

		case "deleted_at.gt":
			queryMods = append(queryMods, qm.Where("deleted_at > ?", v[0]))
		case "deleted_at.lt":
			queryMods = append(queryMods, qm.Where("deleted_at < ?", v[0]))
		case "deleted_at.gte":
			queryMods = append(queryMods, qm.Where("deleted_at >= ?", v[0]))
		case "deleted_at.lte":
			queryMods = append(queryMods, qm.Where("deleted_at <= ?", v[0]))

		}
	}

	count, err := models.Courses(queryMods...).CountG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	limit, offset := api.ExtractLimitOffset(c)

	queryMods = append(queryMods, qm.Limit(limit), qm.Offset(offset))

	if len(orderBy) > 0 {
		queryMods = append(queryMods, qm.OrderBy(strings.Join(orderBy, ", ")))
	} else {
		queryMods = append(queryMods, qm.OrderBy("created_at DESC"))
	}

	courses, err := models.Courses(queryMods...).AllG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	if courses == nil {
		courses = models.CourseSlice{}
	}

	c.JSON(http.StatusOK, GetCoursesResponse{
		Courses:    courses,
		Total:      count,
		NextOffset: int64(offset + limit),
	})
}
