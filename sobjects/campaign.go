// This file was generated for SObject Campaign, API Version v43.0 at 2018-07-30 03:47:34.180752695 -0400 EDT m=+20.524129253

package sobjects

import (
	"fmt"
	"strings"
)

type Campaign struct {
	BaseSObject
	ActualCost                 string  `force:",omitempty"`
	AmountAllOpportunities     string  `force:",omitempty"`
	AmountWonOpportunities     string  `force:",omitempty"`
	BudgetedCost               string  `force:",omitempty"`
	CampaignMemberRecordTypeId string  `force:",omitempty"`
	CreatedById                string  `force:",omitempty"`
	CreatedDate                string  `force:",omitempty"`
	Description                string  `force:",omitempty"`
	EndDate                    string  `force:",omitempty"`
	ExpectedResponse           string  `force:",omitempty"`
	ExpectedRevenue            string  `force:",omitempty"`
	Id                         string  `force:",omitempty"`
	IsActive                   bool    `force:",omitempty"`
	IsDeleted                  bool    `force:",omitempty"`
	LastActivityDate           string  `force:",omitempty"`
	LastModifiedById           string  `force:",omitempty"`
	LastModifiedDate           string  `force:",omitempty"`
	LastReferencedDate         string  `force:",omitempty"`
	LastViewedDate             string  `force:",omitempty"`
	Name                       string  `force:",omitempty"`
	NumberOfContacts           int     `force:",omitempty"`
	NumberOfConvertedLeads     int     `force:",omitempty"`
	NumberOfLeads              int     `force:",omitempty"`
	NumberOfOpportunities      int     `force:",omitempty"`
	NumberOfResponses          int     `force:",omitempty"`
	NumberOfWonOpportunities   int     `force:",omitempty"`
	NumberSent                 float64 `force:",omitempty"`
	OwnerId                    string  `force:",omitempty"`
	ParentId                   string  `force:",omitempty"`
	StartDate                  string  `force:",omitempty"`
	Status                     string  `force:",omitempty"`
	SystemModstamp             string  `force:",omitempty"`
	Type                       string  `force:",omitempty"`
}

func (t *Campaign) ApiName() string {
	return "Campaign"
}

func (t *Campaign) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Campaign #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tActualCost: %v\n", t.ActualCost))
	builder.WriteString(fmt.Sprintf("\tAmountAllOpportunities: %v\n", t.AmountAllOpportunities))
	builder.WriteString(fmt.Sprintf("\tAmountWonOpportunities: %v\n", t.AmountWonOpportunities))
	builder.WriteString(fmt.Sprintf("\tBudgetedCost: %v\n", t.BudgetedCost))
	builder.WriteString(fmt.Sprintf("\tCampaignMemberRecordTypeId: %v\n", t.CampaignMemberRecordTypeId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tEndDate: %v\n", t.EndDate))
	builder.WriteString(fmt.Sprintf("\tExpectedResponse: %v\n", t.ExpectedResponse))
	builder.WriteString(fmt.Sprintf("\tExpectedRevenue: %v\n", t.ExpectedRevenue))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastActivityDate: %v\n", t.LastActivityDate))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNumberOfContacts: %v\n", t.NumberOfContacts))
	builder.WriteString(fmt.Sprintf("\tNumberOfConvertedLeads: %v\n", t.NumberOfConvertedLeads))
	builder.WriteString(fmt.Sprintf("\tNumberOfLeads: %v\n", t.NumberOfLeads))
	builder.WriteString(fmt.Sprintf("\tNumberOfOpportunities: %v\n", t.NumberOfOpportunities))
	builder.WriteString(fmt.Sprintf("\tNumberOfResponses: %v\n", t.NumberOfResponses))
	builder.WriteString(fmt.Sprintf("\tNumberOfWonOpportunities: %v\n", t.NumberOfWonOpportunities))
	builder.WriteString(fmt.Sprintf("\tNumberSent: %v\n", t.NumberSent))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tStartDate: %v\n", t.StartDate))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type CampaignQueryResponse struct {
	BaseQuery
	Records []Campaign `json:"Records" force:"records"`
}
