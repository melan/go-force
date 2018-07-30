// This file was generated for SObject CaseShare, API Version v43.0 at 2018-07-30 03:48:00.375696929 -0400 EDT m=+46.720056425

package sobjects

import (
	"fmt"
	"strings"
)

type CaseShare struct {
	BaseSObject
	CaseAccessLevel  string `force:",omitempty"`
	CaseId           string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	RowCause         string `force:",omitempty"`
	UserOrGroupId    string `force:",omitempty"`
}

func (t *CaseShare) ApiName() string {
	return "CaseShare"
}

func (t *CaseShare) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CaseShare #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCaseAccessLevel: %v\n", t.CaseAccessLevel))
	builder.WriteString(fmt.Sprintf("\tCaseId: %v\n", t.CaseId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tRowCause: %v\n", t.RowCause))
	builder.WriteString(fmt.Sprintf("\tUserOrGroupId: %v\n", t.UserOrGroupId))

	return builder.String()
}

type CaseShareQueryResponse struct {
	BaseQuery
	Records []CaseShare `json:"Records" force:"records"`
}
