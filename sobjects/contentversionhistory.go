// This file was generated for SObject ContentVersionHistory, API Version v43.0 at 2018-07-30 03:47:55.851024954 -0400 EDT m=+42.195214667

package sobjects

import (
	"fmt"
	"strings"
)

type ContentVersionHistory struct {
	BaseSObject
	ContentVersionId string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Field            string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	NewValue         string `force:",omitempty"`
	OldValue         string `force:",omitempty"`
}

func (t *ContentVersionHistory) ApiName() string {
	return "ContentVersionHistory"
}

func (t *ContentVersionHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentVersionHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentVersionId: %v\n", t.ContentVersionId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))

	return builder.String()
}

type ContentVersionHistoryQueryResponse struct {
	BaseQuery
	Records []ContentVersionHistory `json:"Records" force:"records"`
}
