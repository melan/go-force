// This file was generated for SObject CaseTeamRole, API Version v43.0 at 2018-07-30 03:47:57.447844722 -0400 EDT m=+43.792094354

package sobjects

import (
	"fmt"
	"strings"
)

type CaseTeamRole struct {
	BaseSObject
	AccessLevel             string `force:",omitempty"`
	CreatedById             string `force:",omitempty"`
	CreatedDate             string `force:",omitempty"`
	Id                      string `force:",omitempty"`
	LastModifiedById        string `force:",omitempty"`
	LastModifiedDate        string `force:",omitempty"`
	Name                    string `force:",omitempty"`
	PreferencesVisibleInCSP bool   `force:",omitempty"`
	SystemModstamp          string `force:",omitempty"`
}

func (t *CaseTeamRole) ApiName() string {
	return "CaseTeamRole"
}

func (t *CaseTeamRole) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CaseTeamRole #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccessLevel: %v\n", t.AccessLevel))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tPreferencesVisibleInCSP: %v\n", t.PreferencesVisibleInCSP))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CaseTeamRoleQueryResponse struct {
	BaseQuery
	Records []CaseTeamRole `json:"Records" force:"records"`
}
