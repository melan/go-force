// This file was generated for SObject CampaignHistory, API Version v43.0 at 2018-07-30 03:47:23.822460037 -0400 EDT m=+10.165447912

package sobjects

import (
	"fmt"
	"strings"
)

type CampaignHistory struct {
	BaseSObject
	CampaignId  string `force:",omitempty"`
	CreatedById string `force:",omitempty"`
	CreatedDate string `force:",omitempty"`
	Field       string `force:",omitempty"`
	Id          string `force:",omitempty"`
	IsDeleted   bool   `force:",omitempty"`
	NewValue    string `force:",omitempty"`
	OldValue    string `force:",omitempty"`
}

func (t *CampaignHistory) ApiName() string {
	return "CampaignHistory"
}

func (t *CampaignHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CampaignHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCampaignId: %v\n", t.CampaignId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))

	return builder.String()
}

type CampaignHistoryQueryResponse struct {
	BaseQuery
	Records []CampaignHistory `json:"Records" force:"records"`
}
