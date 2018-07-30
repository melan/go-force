// This file was generated for SObject CaseSolution, API Version v43.0 at 2018-07-30 03:47:35.452813234 -0400 EDT m=+21.796237525

package sobjects

import (
	"fmt"
	"strings"
)

type CaseSolution struct {
	BaseSObject
	CaseId         string `force:",omitempty"`
	CreatedById    string `force:",omitempty"`
	CreatedDate    string `force:",omitempty"`
	Id             string `force:",omitempty"`
	IsDeleted      bool   `force:",omitempty"`
	SolutionId     string `force:",omitempty"`
	SystemModstamp string `force:",omitempty"`
}

func (t *CaseSolution) ApiName() string {
	return "CaseSolution"
}

func (t *CaseSolution) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CaseSolution #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCaseId: %v\n", t.CaseId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tSolutionId: %v\n", t.SolutionId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CaseSolutionQueryResponse struct {
	BaseQuery
	Records []CaseSolution `json:"Records" force:"records"`
}
