// This file was generated for SObject Lead, API Version v43.0 at 2018-07-30 03:47:38.089446383 -0400 EDT m=+24.432969612

package sobjects

import (
	"fmt"
	"strings"
)

type Lead struct {
	BaseSObject
	Address                *Address `force:",omitempty"`
	AnnualRevenue          string   `force:",omitempty"`
	City                   string   `force:",omitempty"`
	CleanStatus            string   `force:",omitempty"`
	Company                string   `force:",omitempty"`
	CompanyDunsNumber      string   `force:",omitempty"`
	ConvertedAccountId     string   `force:",omitempty"`
	ConvertedContactId     string   `force:",omitempty"`
	ConvertedDate          string   `force:",omitempty"`
	ConvertedOpportunityId string   `force:",omitempty"`
	Country                string   `force:",omitempty"`
	CreatedById            string   `force:",omitempty"`
	CreatedDate            string   `force:",omitempty"`
	CurrentGenerators__c   string   `force:",omitempty"`
	DandbCompanyId         string   `force:",omitempty"`
	Description            string   `force:",omitempty"`
	Email                  string   `force:",omitempty"`
	EmailBouncedDate       string   `force:",omitempty"`
	EmailBouncedReason     string   `force:",omitempty"`
	Fax                    string   `force:",omitempty"`
	FirstName              string   `force:",omitempty"`
	GeocodeAccuracy        string   `force:",omitempty"`
	Id                     string   `force:",omitempty"`
	Industry               string   `force:",omitempty"`
	IsConverted            bool     `force:",omitempty"`
	IsDeleted              bool     `force:",omitempty"`
	IsUnreadByOwner        bool     `force:",omitempty"`
	Jigsaw                 string   `force:",omitempty"`
	JigsawContactId        string   `force:",omitempty"`
	LastActivityDate       string   `force:",omitempty"`
	LastModifiedById       string   `force:",omitempty"`
	LastModifiedDate       string   `force:",omitempty"`
	LastName               string   `force:",omitempty"`
	LastReferencedDate     string   `force:",omitempty"`
	LastViewedDate         string   `force:",omitempty"`
	Latitude               float64  `force:",omitempty"`
	LeadSource             string   `force:",omitempty"`
	Longitude              float64  `force:",omitempty"`
	MasterRecordId         string   `force:",omitempty"`
	MobilePhone            string   `force:",omitempty"`
	Name                   string   `force:",omitempty"`
	NumberOfEmployees      int      `force:",omitempty"`
	NumberofLocations__c   float64  `force:",omitempty"`
	OwnerId                string   `force:",omitempty"`
	Phone                  string   `force:",omitempty"`
	PhotoUrl               string   `force:",omitempty"`
	PostalCode             string   `force:",omitempty"`
	Primary__c             string   `force:",omitempty"`
	ProductInterest__c     string   `force:",omitempty"`
	Rating                 string   `force:",omitempty"`
	SICCode__c             string   `force:",omitempty"`
	Salutation             string   `force:",omitempty"`
	State                  string   `force:",omitempty"`
	Status                 string   `force:",omitempty"`
	Street                 string   `force:",omitempty"`
	SystemModstamp         string   `force:",omitempty"`
	Title                  string   `force:",omitempty"`
	Website                string   `force:",omitempty"`
}

func (t *Lead) ApiName() string {
	return "Lead"
}

func (t *Lead) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Lead #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAddress: %v\n", t.Address))
	builder.WriteString(fmt.Sprintf("\tAnnualRevenue: %v\n", t.AnnualRevenue))
	builder.WriteString(fmt.Sprintf("\tCity: %v\n", t.City))
	builder.WriteString(fmt.Sprintf("\tCleanStatus: %v\n", t.CleanStatus))
	builder.WriteString(fmt.Sprintf("\tCompany: %v\n", t.Company))
	builder.WriteString(fmt.Sprintf("\tCompanyDunsNumber: %v\n", t.CompanyDunsNumber))
	builder.WriteString(fmt.Sprintf("\tConvertedAccountId: %v\n", t.ConvertedAccountId))
	builder.WriteString(fmt.Sprintf("\tConvertedContactId: %v\n", t.ConvertedContactId))
	builder.WriteString(fmt.Sprintf("\tConvertedDate: %v\n", t.ConvertedDate))
	builder.WriteString(fmt.Sprintf("\tConvertedOpportunityId: %v\n", t.ConvertedOpportunityId))
	builder.WriteString(fmt.Sprintf("\tCountry: %v\n", t.Country))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCurrentGenerators__c: %v\n", t.CurrentGenerators__c))
	builder.WriteString(fmt.Sprintf("\tDandbCompanyId: %v\n", t.DandbCompanyId))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tEmail: %v\n", t.Email))
	builder.WriteString(fmt.Sprintf("\tEmailBouncedDate: %v\n", t.EmailBouncedDate))
	builder.WriteString(fmt.Sprintf("\tEmailBouncedReason: %v\n", t.EmailBouncedReason))
	builder.WriteString(fmt.Sprintf("\tFax: %v\n", t.Fax))
	builder.WriteString(fmt.Sprintf("\tFirstName: %v\n", t.FirstName))
	builder.WriteString(fmt.Sprintf("\tGeocodeAccuracy: %v\n", t.GeocodeAccuracy))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIndustry: %v\n", t.Industry))
	builder.WriteString(fmt.Sprintf("\tIsConverted: %v\n", t.IsConverted))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsUnreadByOwner: %v\n", t.IsUnreadByOwner))
	builder.WriteString(fmt.Sprintf("\tJigsaw: %v\n", t.Jigsaw))
	builder.WriteString(fmt.Sprintf("\tJigsawContactId: %v\n", t.JigsawContactId))
	builder.WriteString(fmt.Sprintf("\tLastActivityDate: %v\n", t.LastActivityDate))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastName: %v\n", t.LastName))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tLatitude: %v\n", t.Latitude))
	builder.WriteString(fmt.Sprintf("\tLeadSource: %v\n", t.LeadSource))
	builder.WriteString(fmt.Sprintf("\tLongitude: %v\n", t.Longitude))
	builder.WriteString(fmt.Sprintf("\tMasterRecordId: %v\n", t.MasterRecordId))
	builder.WriteString(fmt.Sprintf("\tMobilePhone: %v\n", t.MobilePhone))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNumberOfEmployees: %v\n", t.NumberOfEmployees))
	builder.WriteString(fmt.Sprintf("\tNumberofLocations__c: %v\n", t.NumberofLocations__c))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tPhone: %v\n", t.Phone))
	builder.WriteString(fmt.Sprintf("\tPhotoUrl: %v\n", t.PhotoUrl))
	builder.WriteString(fmt.Sprintf("\tPostalCode: %v\n", t.PostalCode))
	builder.WriteString(fmt.Sprintf("\tPrimary__c: %v\n", t.Primary__c))
	builder.WriteString(fmt.Sprintf("\tProductInterest__c: %v\n", t.ProductInterest__c))
	builder.WriteString(fmt.Sprintf("\tRating: %v\n", t.Rating))
	builder.WriteString(fmt.Sprintf("\tSICCode__c: %v\n", t.SICCode__c))
	builder.WriteString(fmt.Sprintf("\tSalutation: %v\n", t.Salutation))
	builder.WriteString(fmt.Sprintf("\tState: %v\n", t.State))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tStreet: %v\n", t.Street))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))
	builder.WriteString(fmt.Sprintf("\tWebsite: %v\n", t.Website))

	return builder.String()
}

type LeadQueryResponse struct {
	BaseQuery
	Records []Lead `json:"Records" force:"records"`
}
