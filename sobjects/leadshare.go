// This file was generated for SObject LeadShare, API Version v43.0 at 2018-07-30 03:47:36.106654431 -0400 EDT m=+22.450103258

package sobjects

import (
	"fmt"
	"strings"
)

type LeadShare struct {
	BaseSObject
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	LeadAccessLevel  string `force:",omitempty"`
	LeadId           string `force:",omitempty"`
	RowCause         string `force:",omitempty"`
	UserOrGroupId    string `force:",omitempty"`
}

func (t *LeadShare) ApiName() string {
	return "LeadShare"
}

func (t *LeadShare) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("LeadShare #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLeadAccessLevel: %v\n", t.LeadAccessLevel))
	builder.WriteString(fmt.Sprintf("\tLeadId: %v\n", t.LeadId))
	builder.WriteString(fmt.Sprintf("\tRowCause: %v\n", t.RowCause))
	builder.WriteString(fmt.Sprintf("\tUserOrGroupId: %v\n", t.UserOrGroupId))

	return builder.String()
}

type LeadShareQueryResponse struct {
	BaseQuery
	Records []LeadShare `json:"Records" force:"records"`
}
