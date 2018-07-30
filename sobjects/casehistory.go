// This file was generated for SObject CaseHistory, API Version v43.0 at 2018-07-30 03:47:16.006514789 -0400 EDT m=+2.349209376

package sobjects

import (
	"fmt"
	"strings"
)

type CaseHistory struct {
	BaseSObject
	CaseId      string `force:",omitempty"`
	CreatedById string `force:",omitempty"`
	CreatedDate string `force:",omitempty"`
	Field       string `force:",omitempty"`
	Id          string `force:",omitempty"`
	IsDeleted   bool   `force:",omitempty"`
	NewValue    string `force:",omitempty"`
	OldValue    string `force:",omitempty"`
}

func (t *CaseHistory) ApiName() string {
	return "CaseHistory"
}

func (t *CaseHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CaseHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCaseId: %v\n", t.CaseId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))

	return builder.String()
}

type CaseHistoryQueryResponse struct {
	BaseQuery
	Records []CaseHistory `json:"Records" force:"records"`
}
