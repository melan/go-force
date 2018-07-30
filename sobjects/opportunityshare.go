// This file was generated for SObject OpportunityShare, API Version v43.0 at 2018-07-30 03:47:34.041765292 -0400 EDT m=+20.385136636

package sobjects

import (
	"fmt"
	"strings"
)

type OpportunityShare struct {
	BaseSObject
	Id                     string `force:",omitempty"`
	IsDeleted              bool   `force:",omitempty"`
	LastModifiedById       string `force:",omitempty"`
	LastModifiedDate       string `force:",omitempty"`
	OpportunityAccessLevel string `force:",omitempty"`
	OpportunityId          string `force:",omitempty"`
	RowCause               string `force:",omitempty"`
	UserOrGroupId          string `force:",omitempty"`
}

func (t *OpportunityShare) ApiName() string {
	return "OpportunityShare"
}

func (t *OpportunityShare) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OpportunityShare #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tOpportunityAccessLevel: %v\n", t.OpportunityAccessLevel))
	builder.WriteString(fmt.Sprintf("\tOpportunityId: %v\n", t.OpportunityId))
	builder.WriteString(fmt.Sprintf("\tRowCause: %v\n", t.RowCause))
	builder.WriteString(fmt.Sprintf("\tUserOrGroupId: %v\n", t.UserOrGroupId))

	return builder.String()
}

type OpportunityShareQueryResponse struct {
	BaseQuery
	Records []OpportunityShare `json:"Records" force:"records"`
}
