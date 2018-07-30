// This file was generated for SObject Partner, API Version v43.0 at 2018-07-30 03:47:34.755226163 -0400 EDT m=+21.098624278

package sobjects

import (
	"fmt"
	"strings"
)

type Partner struct {
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

func (t *Partner) ApiName() string {
	return "Partner"
}

func (t *Partner) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Partner #%s - %s\n", t.Id, t.Name))
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

type PartnerQueryResponse struct {
	BaseQuery
	Records []Partner `json:"Records" force:"records"`
}
