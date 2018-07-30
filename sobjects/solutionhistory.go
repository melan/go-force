// This file was generated for SObject SolutionHistory, API Version v43.0 at 2018-07-30 03:47:29.324289765 -0400 EDT m=+15.667484090

package sobjects

import (
	"fmt"
	"strings"
)

type SolutionHistory struct {
	BaseSObject
	CreatedById string `force:",omitempty"`
	CreatedDate string `force:",omitempty"`
	Field       string `force:",omitempty"`
	Id          string `force:",omitempty"`
	IsDeleted   bool   `force:",omitempty"`
	NewValue    string `force:",omitempty"`
	OldValue    string `force:",omitempty"`
	SolutionId  string `force:",omitempty"`
}

func (t *SolutionHistory) ApiName() string {
	return "SolutionHistory"
}

func (t *SolutionHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("SolutionHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))
	builder.WriteString(fmt.Sprintf("\tSolutionId: %v\n", t.SolutionId))

	return builder.String()
}

type SolutionHistoryQueryResponse struct {
	BaseQuery
	Records []SolutionHistory `json:"Records" force:"records"`
}
