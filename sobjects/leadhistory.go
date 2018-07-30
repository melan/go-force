// This file was generated for SObject LeadHistory, API Version v43.0 at 2018-07-30 03:47:48.828043975 -0400 EDT m=+35.171970158

package sobjects

import (
	"fmt"
	"strings"
)

type LeadHistory struct {
	BaseSObject
	CreatedById string `force:",omitempty"`
	CreatedDate string `force:",omitempty"`
	Field       string `force:",omitempty"`
	Id          string `force:",omitempty"`
	IsDeleted   bool   `force:",omitempty"`
	LeadId      string `force:",omitempty"`
	NewValue    string `force:",omitempty"`
	OldValue    string `force:",omitempty"`
}

func (t *LeadHistory) ApiName() string {
	return "LeadHistory"
}

func (t *LeadHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("LeadHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLeadId: %v\n", t.LeadId))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))

	return builder.String()
}

type LeadHistoryQueryResponse struct {
	BaseQuery
	Records []LeadHistory `json:"Records" force:"records"`
}
