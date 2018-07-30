// This file was generated for SObject PartnerRole, API Version v43.0 at 2018-07-30 03:47:26.198342106 -0400 EDT m=+12.541419133

package sobjects

import (
	"fmt"
	"strings"
)

type PartnerRole struct {
	BaseSObject
	ApiName          string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	MasterLabel      string `force:",omitempty"`
	ReverseRole      string `force:",omitempty"`
	SortOrder        int    `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *PartnerRole) ApiName() string {
	return "PartnerRole"
}

func (t *PartnerRole) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("PartnerRole #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApiName: %v\n", t.ApiName))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tReverseRole: %v\n", t.ReverseRole))
	builder.WriteString(fmt.Sprintf("\tSortOrder: %v\n", t.SortOrder))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type PartnerRoleQueryResponse struct {
	BaseQuery
	Records []PartnerRole `json:"Records" force:"records"`
}
