// This file was generated for SObject Opportunity, API Version v43.0 at 2018-07-30 03:42:13.530979077 -0400 EDT m=+2.966707202

package sobjects

import (
	"fmt"
	"strings"
)

type Opportunity struct {
	BaseSObject
	AccountId                     string  `force:",omitempty"`
	Amount                        string  `force:",omitempty"`
	CampaignId                    string  `force:",omitempty"`
	CloseDate                     string  `force:",omitempty"`
	CreatedById                   string  `force:",omitempty"`
	CreatedDate                   string  `force:",omitempty"`
	CurrentGenerators__c          string  `force:",omitempty"`
	DeliveryInstallationStatus__c string  `force:",omitempty"`
	Description                   string  `force:",omitempty"`
	ExpectedRevenue               string  `force:",omitempty"`
	Fiscal                        string  `force:",omitempty"`
	FiscalQuarter                 int     `force:",omitempty"`
	FiscalYear                    int     `force:",omitempty"`
	ForecastCategory              string  `force:",omitempty"`
	ForecastCategoryName          string  `force:",omitempty"`
	HasOpenActivity               bool    `force:",omitempty"`
	HasOpportunityLineItem        bool    `force:",omitempty"`
	HasOverdueTask                bool    `force:",omitempty"`
	Id                            string  `force:",omitempty"`
	IsClosed                      bool    `force:",omitempty"`
	IsDeleted                     bool    `force:",omitempty"`
	IsPrivate                     bool    `force:",omitempty"`
	IsWon                         bool    `force:",omitempty"`
	LastActivityDate              string  `force:",omitempty"`
	LastModifiedById              string  `force:",omitempty"`
	LastModifiedDate              string  `force:",omitempty"`
	LastReferencedDate            string  `force:",omitempty"`
	LastViewedDate                string  `force:",omitempty"`
	LeadSource                    string  `force:",omitempty"`
	MainCompetitors__c            string  `force:",omitempty"`
	Name                          string  `force:",omitempty"`
	NextStep                      string  `force:",omitempty"`
	OrderNumber__c                string  `force:",omitempty"`
	OwnerId                       string  `force:",omitempty"`
	Pricebook2Id                  string  `force:",omitempty"`
	Probability                   string  `force:",omitempty"`
	StageName                     string  `force:",omitempty"`
	SystemModstamp                string  `force:",omitempty"`
	TotalOpportunityQuantity      float64 `force:",omitempty"`
	TrackingNumber__c             string  `force:",omitempty"`
	Type                          string  `force:",omitempty"`
}

func (t *Opportunity) ApiName() string {
	return "Opportunity"
}

func (t *Opportunity) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Opportunity #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccountId: %v\n", t.AccountId))
	builder.WriteString(fmt.Sprintf("\tAmount: %v\n", t.Amount))
	builder.WriteString(fmt.Sprintf("\tCampaignId: %v\n", t.CampaignId))
	builder.WriteString(fmt.Sprintf("\tCloseDate: %v\n", t.CloseDate))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCurrentGenerators__c: %v\n", t.CurrentGenerators__c))
	builder.WriteString(fmt.Sprintf("\tDeliveryInstallationStatus__c: %v\n", t.DeliveryInstallationStatus__c))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tExpectedRevenue: %v\n", t.ExpectedRevenue))
	builder.WriteString(fmt.Sprintf("\tFiscal: %v\n", t.Fiscal))
	builder.WriteString(fmt.Sprintf("\tFiscalQuarter: %v\n", t.FiscalQuarter))
	builder.WriteString(fmt.Sprintf("\tFiscalYear: %v\n", t.FiscalYear))
	builder.WriteString(fmt.Sprintf("\tForecastCategory: %v\n", t.ForecastCategory))
	builder.WriteString(fmt.Sprintf("\tForecastCategoryName: %v\n", t.ForecastCategoryName))
	builder.WriteString(fmt.Sprintf("\tHasOpenActivity: %v\n", t.HasOpenActivity))
	builder.WriteString(fmt.Sprintf("\tHasOpportunityLineItem: %v\n", t.HasOpportunityLineItem))
	builder.WriteString(fmt.Sprintf("\tHasOverdueTask: %v\n", t.HasOverdueTask))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsClosed: %v\n", t.IsClosed))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsPrivate: %v\n", t.IsPrivate))
	builder.WriteString(fmt.Sprintf("\tIsWon: %v\n", t.IsWon))
	builder.WriteString(fmt.Sprintf("\tLastActivityDate: %v\n", t.LastActivityDate))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tLeadSource: %v\n", t.LeadSource))
	builder.WriteString(fmt.Sprintf("\tMainCompetitors__c: %v\n", t.MainCompetitors__c))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNextStep: %v\n", t.NextStep))
	builder.WriteString(fmt.Sprintf("\tOrderNumber__c: %v\n", t.OrderNumber__c))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tPricebook2Id: %v\n", t.Pricebook2Id))
	builder.WriteString(fmt.Sprintf("\tProbability: %v\n", t.Probability))
	builder.WriteString(fmt.Sprintf("\tStageName: %v\n", t.StageName))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTotalOpportunityQuantity: %v\n", t.TotalOpportunityQuantity))
	builder.WriteString(fmt.Sprintf("\tTrackingNumber__c: %v\n", t.TrackingNumber__c))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type OpportunityQueryResponse struct {
	BaseQuery
	Records []Opportunity `json:"Records" force:"records"`
}
