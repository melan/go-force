// This file was generated for SObject Contract, API Version v43.0 at 2018-07-30 03:47:32.202844694 -0400 EDT m=+18.546147034

package sobjects

import (
	"fmt"
	"strings"
)

type Contract struct {
	BaseSObject
	AccountId              string   `force:",omitempty"`
	ActivatedById          string   `force:",omitempty"`
	ActivatedDate          string   `force:",omitempty"`
	BillingAddress         *Address `force:",omitempty"`
	BillingCity            string   `force:",omitempty"`
	BillingCountry         string   `force:",omitempty"`
	BillingGeocodeAccuracy string   `force:",omitempty"`
	BillingLatitude        float64  `force:",omitempty"`
	BillingLongitude       float64  `force:",omitempty"`
	BillingPostalCode      string   `force:",omitempty"`
	BillingState           string   `force:",omitempty"`
	BillingStreet          string   `force:",omitempty"`
	CompanySignedDate      string   `force:",omitempty"`
	CompanySignedId        string   `force:",omitempty"`
	ContractNumber         string   `force:",omitempty"`
	ContractTerm           int      `force:",omitempty"`
	CreatedById            string   `force:",omitempty"`
	CreatedDate            string   `force:",omitempty"`
	CustomerSignedDate     string   `force:",omitempty"`
	CustomerSignedId       string   `force:",omitempty"`
	CustomerSignedTitle    string   `force:",omitempty"`
	Description            string   `force:",omitempty"`
	EndDate                string   `force:",omitempty"`
	Id                     string   `force:",omitempty"`
	IsDeleted              bool     `force:",omitempty"`
	LastActivityDate       string   `force:",omitempty"`
	LastApprovedDate       string   `force:",omitempty"`
	LastModifiedById       string   `force:",omitempty"`
	LastModifiedDate       string   `force:",omitempty"`
	LastReferencedDate     string   `force:",omitempty"`
	LastViewedDate         string   `force:",omitempty"`
	OwnerExpirationNotice  string   `force:",omitempty"`
	OwnerId                string   `force:",omitempty"`
	Pricebook2Id           string   `force:",omitempty"`
	SpecialTerms           string   `force:",omitempty"`
	StartDate              string   `force:",omitempty"`
	Status                 string   `force:",omitempty"`
	StatusCode             string   `force:",omitempty"`
	SystemModstamp         string   `force:",omitempty"`
}

func (t *Contract) ApiName() string {
	return "Contract"
}

func (t *Contract) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Contract #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccountId: %v\n", t.AccountId))
	builder.WriteString(fmt.Sprintf("\tActivatedById: %v\n", t.ActivatedById))
	builder.WriteString(fmt.Sprintf("\tActivatedDate: %v\n", t.ActivatedDate))
	builder.WriteString(fmt.Sprintf("\tBillingAddress: %v\n", t.BillingAddress))
	builder.WriteString(fmt.Sprintf("\tBillingCity: %v\n", t.BillingCity))
	builder.WriteString(fmt.Sprintf("\tBillingCountry: %v\n", t.BillingCountry))
	builder.WriteString(fmt.Sprintf("\tBillingGeocodeAccuracy: %v\n", t.BillingGeocodeAccuracy))
	builder.WriteString(fmt.Sprintf("\tBillingLatitude: %v\n", t.BillingLatitude))
	builder.WriteString(fmt.Sprintf("\tBillingLongitude: %v\n", t.BillingLongitude))
	builder.WriteString(fmt.Sprintf("\tBillingPostalCode: %v\n", t.BillingPostalCode))
	builder.WriteString(fmt.Sprintf("\tBillingState: %v\n", t.BillingState))
	builder.WriteString(fmt.Sprintf("\tBillingStreet: %v\n", t.BillingStreet))
	builder.WriteString(fmt.Sprintf("\tCompanySignedDate: %v\n", t.CompanySignedDate))
	builder.WriteString(fmt.Sprintf("\tCompanySignedId: %v\n", t.CompanySignedId))
	builder.WriteString(fmt.Sprintf("\tContractNumber: %v\n", t.ContractNumber))
	builder.WriteString(fmt.Sprintf("\tContractTerm: %v\n", t.ContractTerm))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCustomerSignedDate: %v\n", t.CustomerSignedDate))
	builder.WriteString(fmt.Sprintf("\tCustomerSignedId: %v\n", t.CustomerSignedId))
	builder.WriteString(fmt.Sprintf("\tCustomerSignedTitle: %v\n", t.CustomerSignedTitle))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tEndDate: %v\n", t.EndDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastActivityDate: %v\n", t.LastActivityDate))
	builder.WriteString(fmt.Sprintf("\tLastApprovedDate: %v\n", t.LastApprovedDate))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tOwnerExpirationNotice: %v\n", t.OwnerExpirationNotice))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tPricebook2Id: %v\n", t.Pricebook2Id))
	builder.WriteString(fmt.Sprintf("\tSpecialTerms: %v\n", t.SpecialTerms))
	builder.WriteString(fmt.Sprintf("\tStartDate: %v\n", t.StartDate))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tStatusCode: %v\n", t.StatusCode))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ContractQueryResponse struct {
	BaseQuery
	Records []Contract `json:"Records" force:"records"`
}
