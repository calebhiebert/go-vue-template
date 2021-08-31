// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package generated

import (
	convert "github.com/calebhiebert/go-vue-template/convert"
	model "github.com/calebhiebert/go-vue-template/graph/model"
	models "github.com/calebhiebert/go-vue-template/models"
	v8 "github.com/volatiletech/null/v8"
	types "github.com/volatiletech/sqlboiler/v4/types"
)

type ConverterImpl struct{}

func (c *ConverterImpl) ConvertAccessLog(source models.AccessLog) model.AccessLog {
	var modelAccessLog model.AccessLog
	modelAccessLog.ID = source.ID
	modelAccessLog.Path = source.Path
	modelAccessLog.RequestBody = c.nullStringToPString(source.RequestBody)
	modelAccessLog.ResponseBody = c.typesJSONToPString(source.ResponseBody)
	modelAccessLog.ResponseCode = source.ResponseCode
	modelAccessLog.ProcessingDuration = source.ProcessingDuration
	pString := source.RequestMethod
	modelAccessLog.RequestMethod = &pString
	modelAccessLog.UserID = convert.NullStringToString(source.UserID)
	modelAccessLog.IPAddress = source.IPAddress
	modelAccessLog.CreatedAt = convert.NullTimeToInt64(source.CreatedAt)
	return modelAccessLog
}
func (c *ConverterImpl) ConvertAccessLogs(source []models.AccessLog) []model.AccessLog {
	modelAccessLogList := make([]model.AccessLog, len(source))
	for i := 0; i < len(source); i++ {
		modelAccessLogList[i] = c.ConvertAccessLog(source[i])
	}
	return modelAccessLogList
}
func (c *ConverterImpl) ConvertTokenIssuance(source models.TokenIssuance) model.TokenIssuance {
	var modelTokenIssuance model.TokenIssuance
	modelTokenIssuance.ID = source.ID
	modelTokenIssuance.UserID = source.UserID
	modelTokenIssuance.IPAddress = source.IPAddress
	modelTokenIssuance.CreatedAt = convert.NullTimeToInt64(source.CreatedAt)
	return modelTokenIssuance
}
func (c *ConverterImpl) ConvertTokenIssuances(source []models.TokenIssuance) []model.TokenIssuance {
	modelTokenIssuanceList := make([]model.TokenIssuance, len(source))
	for i := 0; i < len(source); i++ {
		modelTokenIssuanceList[i] = c.ConvertTokenIssuance(source[i])
	}
	return modelTokenIssuanceList
}
func (c *ConverterImpl) ConvertUser(source models.User) model.User {
	var modelUser model.User
	modelUser.ID = source.ID
	modelUser.Name = source.Name
	modelUser.Login = c.nullStringToPString(source.Login)
	modelUser.Email = source.Email
	modelUser.Sub = c.nullStringToPString(source.Sub)
	modelUser.Roles = c.typesStringArrayToStringList(source.Roles)
	modelUser.CreatedAt = convert.TimeToInt(source.CreatedAt)
	modelUser.UpdatedAt = convert.TimeToInt(source.UpdatedAt)
	modelUser.DeletedAt = c.nullTimeToPInt(source.DeletedAt)
	return modelUser
}
func (c *ConverterImpl) ConvertUsers(source models.UserSlice) []*model.User {
	pModelUserList := make([]*model.User, len(source))
	for i := 0; i < len(source); i++ {
		var pModelUser *model.User
		if source[i] != nil {
			modelUser := c.ConvertUser(*source[i])
			pModelUser = &modelUser
		}
		pModelUserList[i] = pModelUser
	}
	return pModelUserList
}
func (c *ConverterImpl) nullStringToPString(source v8.String) *string {
	xstring := convert.NullStringToString(source)
	return &xstring
}
func (c *ConverterImpl) nullTimeToPInt(source v8.Time) *int {
	xint := convert.NullTimeToInt64(source)
	return &xint
}
func (c *ConverterImpl) typesJSONToPString(source types.JSON) *string {
	xstring := convert.JSONToString(source)
	return &xstring
}
func (c *ConverterImpl) typesStringArrayToStringList(source types.StringArray) []string {
	stringList := make([]string, len(source))
	for i := 0; i < len(source); i++ {
		stringList[i] = source[i]
	}
	return stringList
}