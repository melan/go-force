// This file was generated for SObject PermissionSetAssignment, API Version v43.0 at 2018-07-30 03:47:20.088738976 -0400 EDT m=+6.431586746

package sobjects

import (
	"fmt"
	"strings"
)

type PermissionSetAssignment struct {
	BaseSObject
	AssigneeId      string `force:",omitempty"`
	Id              string `force:",omitempty"`
	PermissionSetId string `force:",omitempty"`
	SystemModstamp  string `force:",omitempty"`
}

func (t *PermissionSetAssignment) ApiName() string {
	return "PermissionSetAssignment"
}

func (t *PermissionSetAssignment) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("PermissionSetAssignment #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAssigneeId: %v\n", t.AssigneeId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tPermissionSetId: %v\n", t.PermissionSetId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type PermissionSetAssignmentQueryResponse struct {
	BaseQuery
	Records []PermissionSetAssignment `json:"Records" force:"records"`
}
