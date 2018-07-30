// This file was generated for SObject AccountPartner, API Version v43.0 at 2018-07-30 03:48:06.597411078 -0400 EDT m=+52.942004038

package sobjects

import (
	"fmt"
	"strings"
)

type AccountPartner struct {
	BaseSObject
	AccountFromId    string `force:",omitempty"`
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

func (t *AccountPartner) ApiName() string {
	return "AccountPartner"
}

func (t *AccountPartner) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AccountPartner #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccountFromId: %v\n", t.AccountFromId))
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

type AccountPartnerQueryResponse struct {
	BaseQuery
	Records []AccountPartner `json:"Records" force:"records"`
}
