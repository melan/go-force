// This file was generated for SObject ContentDocumentHistory, API Version v43.0 at 2018-07-30 03:47:17.642784027 -0400 EDT m=+3.985540014

package sobjects

import (
	"fmt"
	"strings"
)

type ContentDocumentHistory struct {
	BaseSObject
	ContentDocumentId string `force:",omitempty"`
	CreatedById       string `force:",omitempty"`
	CreatedDate       string `force:",omitempty"`
	Field             string `force:",omitempty"`
	Id                string `force:",omitempty"`
	IsDeleted         bool   `force:",omitempty"`
	NewValue          string `force:",omitempty"`
	OldValue          string `force:",omitempty"`
}

func (t *ContentDocumentHistory) ApiName() string {
	return "ContentDocumentHistory"
}

func (t *ContentDocumentHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentDocumentHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentDocumentId: %v\n", t.ContentDocumentId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))

	return builder.String()
}

type ContentDocumentHistoryQueryResponse struct {
	BaseQuery
	Records []ContentDocumentHistory `json:"Records" force:"records"`
}
