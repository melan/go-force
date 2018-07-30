// This file was generated for SObject CaseContactRole, API Version v43.0 at 2018-07-30 03:47:28.151192914 -0400 EDT m=+14.494343220

package sobjects

import (
	"fmt"
	"strings"
)

type CaseContactRole struct {
	BaseSObject
	CasesId          string `force:",omitempty"`
	ContactId        string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Role             string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *CaseContactRole) ApiName() string {
	return "CaseContactRole"
}

func (t *CaseContactRole) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CaseContactRole #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCasesId: %v\n", t.CasesId))
	builder.WriteString(fmt.Sprintf("\tContactId: %v\n", t.ContactId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tRole: %v\n", t.Role))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CaseContactRoleQueryResponse struct {
	BaseQuery
	Records []CaseContactRole `json:"Records" force:"records"`
}
