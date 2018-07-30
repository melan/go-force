// This file was generated for SObject OpportunityFieldHistory, API Version v43.0 at 2018-07-30 03:48:03.636676646 -0400 EDT m=+49.981158507

package sobjects

import (
	"fmt"
	"strings"
)

type OpportunityFieldHistory struct {
	BaseSObject
	CreatedById   string `force:",omitempty"`
	CreatedDate   string `force:",omitempty"`
	Field         string `force:",omitempty"`
	Id            string `force:",omitempty"`
	IsDeleted     bool   `force:",omitempty"`
	NewValue      string `force:",omitempty"`
	OldValue      string `force:",omitempty"`
	OpportunityId string `force:",omitempty"`
}

func (t *OpportunityFieldHistory) ApiName() string {
	return "OpportunityFieldHistory"
}

func (t *OpportunityFieldHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OpportunityFieldHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))
	builder.WriteString(fmt.Sprintf("\tOpportunityId: %v\n", t.OpportunityId))

	return builder.String()
}

type OpportunityFieldHistoryQueryResponse struct {
	BaseQuery
	Records []OpportunityFieldHistory `json:"Records" force:"records"`
}
