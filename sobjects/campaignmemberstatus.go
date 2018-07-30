// This file was generated for SObject CampaignMemberStatus, API Version v43.0 at 2018-07-30 03:47:48.962031412 -0400 EDT m=+35.305962623

package sobjects

import (
	"fmt"
	"strings"
)

type CampaignMemberStatus struct {
	BaseSObject
	CampaignId       string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	HasResponded     bool   `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDefault        bool   `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	Label            string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	SortOrder        int    `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *CampaignMemberStatus) ApiName() string {
	return "CampaignMemberStatus"
}

func (t *CampaignMemberStatus) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CampaignMemberStatus #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCampaignId: %v\n", t.CampaignId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tHasResponded: %v\n", t.HasResponded))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDefault: %v\n", t.IsDefault))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLabel: %v\n", t.Label))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tSortOrder: %v\n", t.SortOrder))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CampaignMemberStatusQueryResponse struct {
	BaseQuery
	Records []CampaignMemberStatus `json:"Records" force:"records"`
}
