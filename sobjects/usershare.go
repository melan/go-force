// This file was generated for SObject UserShare, API Version v43.0 at 2018-07-30 03:47:24.30264581 -0400 EDT m=+10.645651703

package sobjects

import (
	"fmt"
	"strings"
)

type UserShare struct {
	BaseSObject
	Id               string `force:",omitempty"`
	IsActive         bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	RowCause         string `force:",omitempty"`
	UserAccessLevel  string `force:",omitempty"`
	UserId           string `force:",omitempty"`
	UserOrGroupId    string `force:",omitempty"`
}

func (t *UserShare) ApiName() string {
	return "UserShare"
}

func (t *UserShare) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserShare #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tRowCause: %v\n", t.RowCause))
	builder.WriteString(fmt.Sprintf("\tUserAccessLevel: %v\n", t.UserAccessLevel))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))
	builder.WriteString(fmt.Sprintf("\tUserOrGroupId: %v\n", t.UserOrGroupId))

	return builder.String()
}

type UserShareQueryResponse struct {
	BaseQuery
	Records []UserShare `json:"Records" force:"records"`
}
