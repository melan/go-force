// This file was generated for SObject CampaignShare, API Version v43.0 at 2018-07-30 03:48:05.39080168 -0400 EDT m=+51.735349363

package sobjects

import (
	"fmt"
	"strings"
)

type CampaignShare struct {
	BaseSObject
	CampaignAccessLevel string `force:",omitempty"`
	CampaignId          string `force:",omitempty"`
	Id                  string `force:",omitempty"`
	IsDeleted           bool   `force:",omitempty"`
	LastModifiedById    string `force:",omitempty"`
	LastModifiedDate    string `force:",omitempty"`
	RowCause            string `force:",omitempty"`
	UserOrGroupId       string `force:",omitempty"`
}

func (t *CampaignShare) ApiName() string {
	return "CampaignShare"
}

func (t *CampaignShare) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CampaignShare #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCampaignAccessLevel: %v\n", t.CampaignAccessLevel))
	builder.WriteString(fmt.Sprintf("\tCampaignId: %v\n", t.CampaignId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tRowCause: %v\n", t.RowCause))
	builder.WriteString(fmt.Sprintf("\tUserOrGroupId: %v\n", t.UserOrGroupId))

	return builder.String()
}

type CampaignShareQueryResponse struct {
	BaseQuery
	Records []CampaignShare `json:"Records" force:"records"`
}
