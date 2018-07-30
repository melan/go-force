// This file was generated for SObject CampaignMember, API Version v43.0 at 2018-07-30 03:47:54.367806484 -0400 EDT m=+40.711940541

package sobjects

import (
	"fmt"
	"strings"
)

type CampaignMember struct {
	BaseSObject
	CampaignId           string `force:",omitempty"`
	City                 string `force:",omitempty"`
	CompanyOrAccount     string `force:",omitempty"`
	ContactId            string `force:",omitempty"`
	Country              string `force:",omitempty"`
	CreatedById          string `force:",omitempty"`
	CreatedDate          string `force:",omitempty"`
	Description          string `force:",omitempty"`
	DoNotCall            bool   `force:",omitempty"`
	Email                string `force:",omitempty"`
	Fax                  string `force:",omitempty"`
	FirstName            string `force:",omitempty"`
	FirstRespondedDate   string `force:",omitempty"`
	HasOptedOutOfEmail   bool   `force:",omitempty"`
	HasOptedOutOfFax     bool   `force:",omitempty"`
	HasResponded         bool   `force:",omitempty"`
	Id                   string `force:",omitempty"`
	IsDeleted            bool   `force:",omitempty"`
	LastModifiedById     string `force:",omitempty"`
	LastModifiedDate     string `force:",omitempty"`
	LastName             string `force:",omitempty"`
	LeadId               string `force:",omitempty"`
	LeadOrContactId      string `force:",omitempty"`
	LeadOrContactOwnerId string `force:",omitempty"`
	LeadSource           string `force:",omitempty"`
	MobilePhone          string `force:",omitempty"`
	Name                 string `force:",omitempty"`
	Phone                string `force:",omitempty"`
	PostalCode           string `force:",omitempty"`
	Salutation           string `force:",omitempty"`
	State                string `force:",omitempty"`
	Status               string `force:",omitempty"`
	Street               string `force:",omitempty"`
	SystemModstamp       string `force:",omitempty"`
	Title                string `force:",omitempty"`
	Type                 string `force:",omitempty"`
}

func (t *CampaignMember) ApiName() string {
	return "CampaignMember"
}

func (t *CampaignMember) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CampaignMember #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCampaignId: %v\n", t.CampaignId))
	builder.WriteString(fmt.Sprintf("\tCity: %v\n", t.City))
	builder.WriteString(fmt.Sprintf("\tCompanyOrAccount: %v\n", t.CompanyOrAccount))
	builder.WriteString(fmt.Sprintf("\tContactId: %v\n", t.ContactId))
	builder.WriteString(fmt.Sprintf("\tCountry: %v\n", t.Country))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDoNotCall: %v\n", t.DoNotCall))
	builder.WriteString(fmt.Sprintf("\tEmail: %v\n", t.Email))
	builder.WriteString(fmt.Sprintf("\tFax: %v\n", t.Fax))
	builder.WriteString(fmt.Sprintf("\tFirstName: %v\n", t.FirstName))
	builder.WriteString(fmt.Sprintf("\tFirstRespondedDate: %v\n", t.FirstRespondedDate))
	builder.WriteString(fmt.Sprintf("\tHasOptedOutOfEmail: %v\n", t.HasOptedOutOfEmail))
	builder.WriteString(fmt.Sprintf("\tHasOptedOutOfFax: %v\n", t.HasOptedOutOfFax))
	builder.WriteString(fmt.Sprintf("\tHasResponded: %v\n", t.HasResponded))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastName: %v\n", t.LastName))
	builder.WriteString(fmt.Sprintf("\tLeadId: %v\n", t.LeadId))
	builder.WriteString(fmt.Sprintf("\tLeadOrContactId: %v\n", t.LeadOrContactId))
	builder.WriteString(fmt.Sprintf("\tLeadOrContactOwnerId: %v\n", t.LeadOrContactOwnerId))
	builder.WriteString(fmt.Sprintf("\tLeadSource: %v\n", t.LeadSource))
	builder.WriteString(fmt.Sprintf("\tMobilePhone: %v\n", t.MobilePhone))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tPhone: %v\n", t.Phone))
	builder.WriteString(fmt.Sprintf("\tPostalCode: %v\n", t.PostalCode))
	builder.WriteString(fmt.Sprintf("\tSalutation: %v\n", t.Salutation))
	builder.WriteString(fmt.Sprintf("\tState: %v\n", t.State))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tStreet: %v\n", t.Street))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type CampaignMemberQueryResponse struct {
	BaseQuery
	Records []CampaignMember `json:"Records" force:"records"`
}
