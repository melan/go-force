// This file was generated for SObject OpportunityContactRole, API Version v43.0 at 2018-07-30 03:47:44.896697486 -0400 EDT m=+31.240476150

package sobjects

import (
	"fmt"
	"strings"
)

type OpportunityContactRole struct {
	BaseSObject
	ContactId        string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	IsPrimary        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	OpportunityId    string `force:",omitempty"`
	Role             string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *OpportunityContactRole) ApiName() string {
	return "OpportunityContactRole"
}

func (t *OpportunityContactRole) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OpportunityContactRole #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContactId: %v\n", t.ContactId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsPrimary: %v\n", t.IsPrimary))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tOpportunityId: %v\n", t.OpportunityId))
	builder.WriteString(fmt.Sprintf("\tRole: %v\n", t.Role))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type OpportunityContactRoleQueryResponse struct {
	BaseQuery
	Records []OpportunityContactRole `json:"Records" force:"records"`
}
