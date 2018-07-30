// This file was generated for SObject UserAppMenuCustomizationShare, API Version v43.0 at 2018-07-30 03:47:52.512854967 -0400 EDT m=+38.856919419

package sobjects

import (
	"fmt"
	"strings"
)

type UserAppMenuCustomizationShare struct {
	BaseSObject
	AccessLevel      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	ParentId         string `force:",omitempty"`
	RowCause         string `force:",omitempty"`
	UserOrGroupId    string `force:",omitempty"`
}

func (t *UserAppMenuCustomizationShare) ApiName() string {
	return "UserAppMenuCustomizationShare"
}

func (t *UserAppMenuCustomizationShare) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserAppMenuCustomizationShare #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccessLevel: %v\n", t.AccessLevel))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tRowCause: %v\n", t.RowCause))
	builder.WriteString(fmt.Sprintf("\tUserOrGroupId: %v\n", t.UserOrGroupId))

	return builder.String()
}

type UserAppMenuCustomizationShareQueryResponse struct {
	BaseQuery
	Records []UserAppMenuCustomizationShare `json:"Records" force:"records"`
}
