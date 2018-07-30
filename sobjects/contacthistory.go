// This file was generated for SObject ContactHistory, API Version v43.0 at 2018-07-30 03:47:53.834448049 -0400 EDT m=+40.178562092

package sobjects

import (
	"fmt"
	"strings"
)

type ContactHistory struct {
	BaseSObject
	ContactId   string `force:",omitempty"`
	CreatedById string `force:",omitempty"`
	CreatedDate string `force:",omitempty"`
	Field       string `force:",omitempty"`
	Id          string `force:",omitempty"`
	IsDeleted   bool   `force:",omitempty"`
	NewValue    string `force:",omitempty"`
	OldValue    string `force:",omitempty"`
}

func (t *ContactHistory) ApiName() string {
	return "ContactHistory"
}

func (t *ContactHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContactHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContactId: %v\n", t.ContactId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))

	return builder.String()
}

type ContactHistoryQueryResponse struct {
	BaseQuery
	Records []ContactHistory `json:"Records" force:"records"`
}
