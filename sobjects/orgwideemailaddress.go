// This file was generated for SObject OrgWideEmailAddress, API Version v43.0 at 2018-07-30 03:48:10.653907884 -0400 EDT m=+56.998653060

package sobjects

import (
	"fmt"
	"strings"
)

type OrgWideEmailAddress struct {
	BaseSObject
	Address            string `force:",omitempty"`
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	DisplayName        string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsAllowAllProfiles bool   `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
}

func (t *OrgWideEmailAddress) ApiName() string {
	return "OrgWideEmailAddress"
}

func (t *OrgWideEmailAddress) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OrgWideEmailAddress #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAddress: %v\n", t.Address))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDisplayName: %v\n", t.DisplayName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsAllowAllProfiles: %v\n", t.IsAllowAllProfiles))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type OrgWideEmailAddressQueryResponse struct {
	BaseQuery
	Records []OrgWideEmailAddress `json:"Records" force:"records"`
}
