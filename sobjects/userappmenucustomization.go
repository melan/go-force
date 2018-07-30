// This file was generated for SObject UserAppMenuCustomization, API Version v43.0 at 2018-07-30 03:47:23.232941088 -0400 EDT m=+9.575906841

package sobjects

import (
	"fmt"
	"strings"
)

type UserAppMenuCustomization struct {
	BaseSObject
	ApplicationId    string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	OwnerId          string `force:",omitempty"`
	SortOrder        int    `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *UserAppMenuCustomization) ApiName() string {
	return "UserAppMenuCustomization"
}

func (t *UserAppMenuCustomization) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserAppMenuCustomization #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApplicationId: %v\n", t.ApplicationId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tSortOrder: %v\n", t.SortOrder))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type UserAppMenuCustomizationQueryResponse struct {
	BaseQuery
	Records []UserAppMenuCustomization `json:"Records" force:"records"`
}
