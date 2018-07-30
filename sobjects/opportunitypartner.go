// This file was generated for SObject OpportunityPartner, API Version v43.0 at 2018-07-30 03:48:01.216978042 -0400 EDT m=+47.561369106

package sobjects

import (
	"fmt"
	"strings"
)

type OpportunityPartner struct {
	BaseSObject
	AccountToId      string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	IsPrimary        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	OpportunityId    string `force:",omitempty"`
	ReversePartnerId string `force:",omitempty"`
	Role             string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *OpportunityPartner) ApiName() string {
	return "OpportunityPartner"
}

func (t *OpportunityPartner) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OpportunityPartner #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccountToId: %v\n", t.AccountToId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsPrimary: %v\n", t.IsPrimary))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tOpportunityId: %v\n", t.OpportunityId))
	builder.WriteString(fmt.Sprintf("\tReversePartnerId: %v\n", t.ReversePartnerId))
	builder.WriteString(fmt.Sprintf("\tRole: %v\n", t.Role))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type OpportunityPartnerQueryResponse struct {
	BaseQuery
	Records []OpportunityPartner `json:"Records" force:"records"`
}
