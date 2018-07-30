// This file was generated for SObject Contact, API Version v43.0 at 2018-07-30 03:47:52.810496826 -0400 EDT m=+39.154572447

package sobjects

import (
	"fmt"
	"strings"
)

type Contact struct {
	BaseSObject
	AccountId              string   `force:",omitempty"`
	AssistantName          string   `force:",omitempty"`
	AssistantPhone         string   `force:",omitempty"`
	Birthdate              string   `force:",omitempty"`
	CleanStatus            string   `force:",omitempty"`
	CreatedById            string   `force:",omitempty"`
	CreatedDate            string   `force:",omitempty"`
	Department             string   `force:",omitempty"`
	Description            string   `force:",omitempty"`
	Email                  string   `force:",omitempty"`
	EmailBouncedDate       string   `force:",omitempty"`
	EmailBouncedReason     string   `force:",omitempty"`
	Fax                    string   `force:",omitempty"`
	FirstName              string   `force:",omitempty"`
	HomePhone              string   `force:",omitempty"`
	Id                     string   `force:",omitempty"`
	IsDeleted              bool     `force:",omitempty"`
	IsEmailBounced         bool     `force:",omitempty"`
	Jigsaw                 string   `force:",omitempty"`
	JigsawContactId        string   `force:",omitempty"`
	Languages__c           string   `force:",omitempty"`
	LastActivityDate       string   `force:",omitempty"`
	LastCURequestDate      string   `force:",omitempty"`
	LastCUUpdateDate       string   `force:",omitempty"`
	LastModifiedById       string   `force:",omitempty"`
	LastModifiedDate       string   `force:",omitempty"`
	LastName               string   `force:",omitempty"`
	LastReferencedDate     string   `force:",omitempty"`
	LastViewedDate         string   `force:",omitempty"`
	LeadSource             string   `force:",omitempty"`
	Level__c               string   `force:",omitempty"`
	MailingAddress         *Address `force:",omitempty"`
	MailingCity            string   `force:",omitempty"`
	MailingCountry         string   `force:",omitempty"`
	MailingGeocodeAccuracy string   `force:",omitempty"`
	MailingLatitude        float64  `force:",omitempty"`
	MailingLongitude       float64  `force:",omitempty"`
	MailingPostalCode      string   `force:",omitempty"`
	MailingState           string   `force:",omitempty"`
	MailingStreet          string   `force:",omitempty"`
	MasterRecordId         string   `force:",omitempty"`
	MobilePhone            string   `force:",omitempty"`
	Name                   string   `force:",omitempty"`
	OtherAddress           *Address `force:",omitempty"`
	OtherCity              string   `force:",omitempty"`
	OtherCountry           string   `force:",omitempty"`
	OtherGeocodeAccuracy   string   `force:",omitempty"`
	OtherLatitude          float64  `force:",omitempty"`
	OtherLongitude         float64  `force:",omitempty"`
	OtherPhone             string   `force:",omitempty"`
	OtherPostalCode        string   `force:",omitempty"`
	OtherState             string   `force:",omitempty"`
	OtherStreet            string   `force:",omitempty"`
	OwnerId                string   `force:",omitempty"`
	Phone                  string   `force:",omitempty"`
	PhotoUrl               string   `force:",omitempty"`
	ReportsToId            string   `force:",omitempty"`
	Salutation             string   `force:",omitempty"`
	SystemModstamp         string   `force:",omitempty"`
	Title                  string   `force:",omitempty"`
}

func (t *Contact) ApiName() string {
	return "Contact"
}

func (t *Contact) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Contact #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccountId: %v\n", t.AccountId))
	builder.WriteString(fmt.Sprintf("\tAssistantName: %v\n", t.AssistantName))
	builder.WriteString(fmt.Sprintf("\tAssistantPhone: %v\n", t.AssistantPhone))
	builder.WriteString(fmt.Sprintf("\tBirthdate: %v\n", t.Birthdate))
	builder.WriteString(fmt.Sprintf("\tCleanStatus: %v\n", t.CleanStatus))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDepartment: %v\n", t.Department))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tEmail: %v\n", t.Email))
	builder.WriteString(fmt.Sprintf("\tEmailBouncedDate: %v\n", t.EmailBouncedDate))
	builder.WriteString(fmt.Sprintf("\tEmailBouncedReason: %v\n", t.EmailBouncedReason))
	builder.WriteString(fmt.Sprintf("\tFax: %v\n", t.Fax))
	builder.WriteString(fmt.Sprintf("\tFirstName: %v\n", t.FirstName))
	builder.WriteString(fmt.Sprintf("\tHomePhone: %v\n", t.HomePhone))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsEmailBounced: %v\n", t.IsEmailBounced))
	builder.WriteString(fmt.Sprintf("\tJigsaw: %v\n", t.Jigsaw))
	builder.WriteString(fmt.Sprintf("\tJigsawContactId: %v\n", t.JigsawContactId))
	builder.WriteString(fmt.Sprintf("\tLanguages__c: %v\n", t.Languages__c))
	builder.WriteString(fmt.Sprintf("\tLastActivityDate: %v\n", t.LastActivityDate))
	builder.WriteString(fmt.Sprintf("\tLastCURequestDate: %v\n", t.LastCURequestDate))
	builder.WriteString(fmt.Sprintf("\tLastCUUpdateDate: %v\n", t.LastCUUpdateDate))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastName: %v\n", t.LastName))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tLeadSource: %v\n", t.LeadSource))
	builder.WriteString(fmt.Sprintf("\tLevel__c: %v\n", t.Level__c))
	builder.WriteString(fmt.Sprintf("\tMailingAddress: %v\n", t.MailingAddress))
	builder.WriteString(fmt.Sprintf("\tMailingCity: %v\n", t.MailingCity))
	builder.WriteString(fmt.Sprintf("\tMailingCountry: %v\n", t.MailingCountry))
	builder.WriteString(fmt.Sprintf("\tMailingGeocodeAccuracy: %v\n", t.MailingGeocodeAccuracy))
	builder.WriteString(fmt.Sprintf("\tMailingLatitude: %v\n", t.MailingLatitude))
	builder.WriteString(fmt.Sprintf("\tMailingLongitude: %v\n", t.MailingLongitude))
	builder.WriteString(fmt.Sprintf("\tMailingPostalCode: %v\n", t.MailingPostalCode))
	builder.WriteString(fmt.Sprintf("\tMailingState: %v\n", t.MailingState))
	builder.WriteString(fmt.Sprintf("\tMailingStreet: %v\n", t.MailingStreet))
	builder.WriteString(fmt.Sprintf("\tMasterRecordId: %v\n", t.MasterRecordId))
	builder.WriteString(fmt.Sprintf("\tMobilePhone: %v\n", t.MobilePhone))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOtherAddress: %v\n", t.OtherAddress))
	builder.WriteString(fmt.Sprintf("\tOtherCity: %v\n", t.OtherCity))
	builder.WriteString(fmt.Sprintf("\tOtherCountry: %v\n", t.OtherCountry))
	builder.WriteString(fmt.Sprintf("\tOtherGeocodeAccuracy: %v\n", t.OtherGeocodeAccuracy))
	builder.WriteString(fmt.Sprintf("\tOtherLatitude: %v\n", t.OtherLatitude))
	builder.WriteString(fmt.Sprintf("\tOtherLongitude: %v\n", t.OtherLongitude))
	builder.WriteString(fmt.Sprintf("\tOtherPhone: %v\n", t.OtherPhone))
	builder.WriteString(fmt.Sprintf("\tOtherPostalCode: %v\n", t.OtherPostalCode))
	builder.WriteString(fmt.Sprintf("\tOtherState: %v\n", t.OtherState))
	builder.WriteString(fmt.Sprintf("\tOtherStreet: %v\n", t.OtherStreet))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tPhone: %v\n", t.Phone))
	builder.WriteString(fmt.Sprintf("\tPhotoUrl: %v\n", t.PhotoUrl))
	builder.WriteString(fmt.Sprintf("\tReportsToId: %v\n", t.ReportsToId))
	builder.WriteString(fmt.Sprintf("\tSalutation: %v\n", t.Salutation))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))

	return builder.String()
}

type ContactQueryResponse struct {
	BaseQuery
	Records []Contact `json:"Records" force:"records"`
}
