// This file was generated for SObject UserListView, API Version v43.0 at 2018-07-30 03:47:43.384009417 -0400 EDT m=+29.727731318

package sobjects

import (
	"fmt"
	"strings"
)

type UserListView struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	LastViewedChart  string `force:",omitempty"`
	ListViewId       string `force:",omitempty"`
	SobjectType      string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	UserId           string `force:",omitempty"`
}

func (t *UserListView) ApiName() string {
	return "UserListView"
}

func (t *UserListView) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserListView #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedChart: %v\n", t.LastViewedChart))
	builder.WriteString(fmt.Sprintf("\tListViewId: %v\n", t.ListViewId))
	builder.WriteString(fmt.Sprintf("\tSobjectType: %v\n", t.SobjectType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type UserListViewQueryResponse struct {
	BaseQuery
	Records []UserListView `json:"Records" force:"records"`
}
