// This file was generated for SObject ListEmailRecipientSource, API Version v43.0 at 2018-07-30 03:47:39.826085615 -0400 EDT m=+26.169674009

package sobjects

import (
	"fmt"
	"strings"
)

type ListEmailRecipientSource struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	ListEmailId      string `force:",omitempty"`
	Name             string `force:",omitempty"`
	SourceListId     string `force:",omitempty"`
	SourceType       string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *ListEmailRecipientSource) ApiName() string {
	return "ListEmailRecipientSource"
}

func (t *ListEmailRecipientSource) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ListEmailRecipientSource #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tListEmailId: %v\n", t.ListEmailId))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tSourceListId: %v\n", t.SourceListId))
	builder.WriteString(fmt.Sprintf("\tSourceType: %v\n", t.SourceType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ListEmailRecipientSourceQueryResponse struct {
	BaseQuery
	Records []ListEmailRecipientSource `json:"Records" force:"records"`
}
