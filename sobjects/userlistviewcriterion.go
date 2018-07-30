// This file was generated for SObject UserListViewCriterion, API Version v43.0 at 2018-07-30 03:47:30.778534462 -0400 EDT m=+17.121783356

package sobjects

import (
	"fmt"
	"strings"
)

type UserListViewCriterion struct {
	BaseSObject
	ColumnName       string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Operation        string `force:",omitempty"`
	SortOrder        int    `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	UserListViewId   string `force:",omitempty"`
	Value            string `force:",omitempty"`
}

func (t *UserListViewCriterion) ApiName() string {
	return "UserListViewCriterion"
}

func (t *UserListViewCriterion) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserListViewCriterion #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tColumnName: %v\n", t.ColumnName))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tOperation: %v\n", t.Operation))
	builder.WriteString(fmt.Sprintf("\tSortOrder: %v\n", t.SortOrder))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserListViewId: %v\n", t.UserListViewId))
	builder.WriteString(fmt.Sprintf("\tValue: %v\n", t.Value))

	return builder.String()
}

type UserListViewCriterionQueryResponse struct {
	BaseQuery
	Records []UserListViewCriterion `json:"Records" force:"records"`
}
