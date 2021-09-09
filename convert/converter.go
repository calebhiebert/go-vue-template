package convert

import (
	"time"

	"github.com/calebhiebert/go-vue-template/graph/model"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// goverter:converter
// goverter:extend NullStringToString
// goverter:extend JSONToString
// goverter:extend NullTimeToString
// goverter:extend TimeToString
type Converter interface {
	// goverter:ignore AvatarURL
	ConvertUser(input models.User) model.User

	// goverter:ignore AvatarURL
	ConvertUsers(input models.UserSlice) []*model.User

	// goverter:ignore User
	ConvertTokenIssuance(input models.TokenIssuance) model.TokenIssuance

	// goverter:ignore User
	ConvertTokenIssuances(input []models.TokenIssuance) []model.TokenIssuance

	// goverter:ignore User
	ConvertAccessLog(input models.AccessLog) model.AccessLog

	// goverter:ignore User
	ConvertAccessLogs(input []models.AccessLog) []model.AccessLog
}

func NullStringToString(str null.String) string {
	return str.String
}

func JSONToString(json types.JSON) string {
	return json.String()
}

func NullTimeToInt64(time null.Time) int {
	return int(time.Time.Unix())
}

func TimeToInt(time time.Time) int {
	return int(time.Unix())
}

func NullTimeToString(t null.Time) string {
	return t.Time.Format(time.RFC3339)
}

func TimeToString(t time.Time) string {
	return t.Format(time.RFC3339)
}
