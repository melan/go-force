// This file was generated for SObject CategoryNode, API Version v43.0 at 2018-07-30 03:48:03.276575124 -0400 EDT m=+49.621043473

package sobjects

import (
	"fmt"
	"strings"
)

type CategoryNode struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	MasterLabel      string `force:",omitempty"`
	ParentId         string `force:",omitempty"`
	SortOrder        int    `force:",omitempty"`
	SortStyle        string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *CategoryNode) ApiName() string {
	return "CategoryNode"
}

func (t *CategoryNode) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CategoryNode #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tSortOrder: %v\n", t.SortOrder))
	builder.WriteString(fmt.Sprintf("\tSortStyle: %v\n", t.SortStyle))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CategoryNodeQueryResponse struct {
	BaseQuery
	Records []CategoryNode `json:"Records" force:"records"`
}
