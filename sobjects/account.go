// This file was generated for SObject Account, API Version v43.0 at 2018-07-30 03:42:14.277253093 -0400 EDT m=+3.713009221

package sobjects

import (
	"fmt"
	"strings"
)

type Account struct {
	BaseSObject
	AccountNumber           string   `force:",omitempty"`
	AccountSource           string   `force:",omitempty"`
	Active__c               string   `force:",omitempty"`
	AnnualRevenue           string   `force:",omitempty"`
	BillingAddress          *Address `force:",omitempty"`
	BillingCity             string   `force:",omitempty"`
	BillingCountry          string   `force:",omitempty"`
	BillingGeocodeAccuracy  string   `force:",omitempty"`
	BillingLatitude         float64  `force:",omitempty"`
	BillingLongitude        float64  `force:",omitempty"`
	BillingPostalCode       string   `force:",omitempty"`
	BillingState            string   `force:",omitempty"`
	BillingStreet           string   `force:",omitempty"`
	CleanStatus             string   `force:",omitempty"`
	CreatedById             string   `force:",omitempty"`
	CreatedDate             string   `force:",omitempty"`
	CustomerPriority__c     string   `force:",omitempty"`
	DandbCompanyId          string   `force:",omitempty"`
	Description             string   `force:",omitempty"`
	DunsNumber              string   `force:",omitempty"`
	Fax                     string   `force:",omitempty"`
	Id                      string   `force:",omitempty"`
	Industry                string   `force:",omitempty"`
	IsDeleted               bool     `force:",omitempty"`
	Jigsaw                  string   `force:",omitempty"`
	JigsawCompanyId         string   `force:",omitempty"`
	LastActivityDate        string   `force:",omitempty"`
	LastModifiedById        string   `force:",omitempty"`
	LastModifiedDate        string   `force:",omitempty"`
	LastReferencedDate      string   `force:",omitempty"`
	LastViewedDate          string   `force:",omitempty"`
	MasterRecordId          string   `force:",omitempty"`
	NaicsCode               string   `force:",omitempty"`
	NaicsDesc               string   `force:",omitempty"`
	Name                    string   `force:",omitempty"`
	NumberOfEmployees       int      `force:",omitempty"`
	NumberofLocations__c    float64  `force:",omitempty"`
	OwnerId                 string   `force:",omitempty"`
	Ownership               string   `force:",omitempty"`
	ParentId                string   `force:",omitempty"`
	Phone                   string   `force:",omitempty"`
	PhotoUrl                string   `force:",omitempty"`
	Rating                  string   `force:",omitempty"`
	SLAExpirationDate__c    string   `force:",omitempty"`
	SLASerialNumber__c      string   `force:",omitempty"`
	SLA__c                  string   `force:",omitempty"`
	ShippingAddress         *Address `force:",omitempty"`
	ShippingCity            string   `force:",omitempty"`
	ShippingCountry         string   `force:",omitempty"`
	ShippingGeocodeAccuracy string   `force:",omitempty"`
	ShippingLatitude        float64  `force:",omitempty"`
	ShippingLongitude       float64  `force:",omitempty"`
	ShippingPostalCode      string   `force:",omitempty"`
	ShippingState           string   `force:",omitempty"`
	ShippingStreet          string   `force:",omitempty"`
	Sic                     string   `force:",omitempty"`
	SicDesc                 string   `force:",omitempty"`
	Site                    string   `force:",omitempty"`
	SystemModstamp          string   `force:",omitempty"`
	TickerSymbol            string   `force:",omitempty"`
	Tradestyle              string   `force:",omitempty"`
	Type                    string   `force:",omitempty"`
	UpsellOpportunity__c    string   `force:",omitempty"`
	Website                 string   `force:",omitempty"`
	YearStarted             string   `force:",omitempty"`
}

func (t *Account) ApiName() string {
	return "Account"
}

func (t *Account) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Account #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccountNumber: %v\n", t.AccountNumber))
	builder.WriteString(fmt.Sprintf("\tAccountSource: %v\n", t.AccountSource))
	builder.WriteString(fmt.Sprintf("\tActive__c: %v\n", t.Active__c))
	builder.WriteString(fmt.Sprintf("\tAnnualRevenue: %v\n", t.AnnualRevenue))
	builder.WriteString(fmt.Sprintf("\tBillingAddress: %v\n", t.BillingAddress))
	builder.WriteString(fmt.Sprintf("\tBillingCity: %v\n", t.BillingCity))
	builder.WriteString(fmt.Sprintf("\tBillingCountry: %v\n", t.BillingCountry))
	builder.WriteString(fmt.Sprintf("\tBillingGeocodeAccuracy: %v\n", t.BillingGeocodeAccuracy))
	builder.WriteString(fmt.Sprintf("\tBillingLatitude: %v\n", t.BillingLatitude))
	builder.WriteString(fmt.Sprintf("\tBillingLongitude: %v\n", t.BillingLongitude))
	builder.WriteString(fmt.Sprintf("\tBillingPostalCode: %v\n", t.BillingPostalCode))
	builder.WriteString(fmt.Sprintf("\tBillingState: %v\n", t.BillingState))
	builder.WriteString(fmt.Sprintf("\tBillingStreet: %v\n", t.BillingStreet))
	builder.WriteString(fmt.Sprintf("\tCleanStatus: %v\n", t.CleanStatus))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCustomerPriority__c: %v\n", t.CustomerPriority__c))
	builder.WriteString(fmt.Sprintf("\tDandbCompanyId: %v\n", t.DandbCompanyId))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDunsNumber: %v\n", t.DunsNumber))
	builder.WriteString(fmt.Sprintf("\tFax: %v\n", t.Fax))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIndustry: %v\n", t.Industry))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tJigsaw: %v\n", t.Jigsaw))
	builder.WriteString(fmt.Sprintf("\tJigsawCompanyId: %v\n", t.JigsawCompanyId))
	builder.WriteString(fmt.Sprintf("\tLastActivityDate: %v\n", t.LastActivityDate))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tMasterRecordId: %v\n", t.MasterRecordId))
	builder.WriteString(fmt.Sprintf("\tNaicsCode: %v\n", t.NaicsCode))
	builder.WriteString(fmt.Sprintf("\tNaicsDesc: %v\n", t.NaicsDesc))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNumberOfEmployees: %v\n", t.NumberOfEmployees))
	builder.WriteString(fmt.Sprintf("\tNumberofLocations__c: %v\n", t.NumberofLocations__c))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tOwnership: %v\n", t.Ownership))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tPhone: %v\n", t.Phone))
	builder.WriteString(fmt.Sprintf("\tPhotoUrl: %v\n", t.PhotoUrl))
	builder.WriteString(fmt.Sprintf("\tRating: %v\n", t.Rating))
	builder.WriteString(fmt.Sprintf("\tSLAExpirationDate__c: %v\n", t.SLAExpirationDate__c))
	builder.WriteString(fmt.Sprintf("\tSLASerialNumber__c: %v\n", t.SLASerialNumber__c))
	builder.WriteString(fmt.Sprintf("\tSLA__c: %v\n", t.SLA__c))
	builder.WriteString(fmt.Sprintf("\tShippingAddress: %v\n", t.ShippingAddress))
	builder.WriteString(fmt.Sprintf("\tShippingCity: %v\n", t.ShippingCity))
	builder.WriteString(fmt.Sprintf("\tShippingCountry: %v\n", t.ShippingCountry))
	builder.WriteString(fmt.Sprintf("\tShippingGeocodeAccuracy: %v\n", t.ShippingGeocodeAccuracy))
	builder.WriteString(fmt.Sprintf("\tShippingLatitude: %v\n", t.ShippingLatitude))
	builder.WriteString(fmt.Sprintf("\tShippingLongitude: %v\n", t.ShippingLongitude))
	builder.WriteString(fmt.Sprintf("\tShippingPostalCode: %v\n", t.ShippingPostalCode))
	builder.WriteString(fmt.Sprintf("\tShippingState: %v\n", t.ShippingState))
	builder.WriteString(fmt.Sprintf("\tShippingStreet: %v\n", t.ShippingStreet))
	builder.WriteString(fmt.Sprintf("\tSic: %v\n", t.Sic))
	builder.WriteString(fmt.Sprintf("\tSicDesc: %v\n", t.SicDesc))
	builder.WriteString(fmt.Sprintf("\tSite: %v\n", t.Site))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTickerSymbol: %v\n", t.TickerSymbol))
	builder.WriteString(fmt.Sprintf("\tTradestyle: %v\n", t.Tradestyle))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))
	builder.WriteString(fmt.Sprintf("\tUpsellOpportunity__c: %v\n", t.UpsellOpportunity__c))
	builder.WriteString(fmt.Sprintf("\tWebsite: %v\n", t.Website))
	builder.WriteString(fmt.Sprintf("\tYearStarted: %v\n", t.YearStarted))

	return builder.String()
}

type AccountQueryResponse struct {
	BaseQuery
	Records []Account `json:"Records" force:"records"`
}
