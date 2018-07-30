// This file was generated for SObject UserAppInfo, API Version v43.0 at 2018-07-30 03:48:09.90337132 -0400 EDT m=+56.248088333

package sobjects

import (
	"fmt"
	"strings"
)

type UserAppInfo struct {
	BaseSObject
	AppDefinitionId  string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	FormFactor       string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	UserId           string `force:",omitempty"`
}

func (t *UserAppInfo) ApiName() string {
	return "UserAppInfo"
}

func (t *UserAppInfo) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserAppInfo #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAppDefinitionId: %v\n", t.AppDefinitionId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFormFactor: %v\n", t.FormFactor))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type UserAppInfoQueryResponse struct {
	BaseQuery
	Records []UserAppInfo `json:"Records" force:"records"`
}
