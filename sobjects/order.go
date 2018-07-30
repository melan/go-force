// This file was generated for SObject Order, API Version v43.0 at 2018-07-30 03:47:56.801421065 -0400 EDT m=+43.145646441

package sobjects

import (
	"fmt"
	"strings"
)

type Order struct {
	BaseSObject
	AccountId               string   `force:",omitempty"`
	ActivatedById           string   `force:",omitempty"`
	ActivatedDate           string   `force:",omitempty"`
	BillToContactId         string   `force:",omitempty"`
	BillingAddress          *Address `force:",omitempty"`
	BillingCity             string   `force:",omitempty"`
	BillingCountry          string   `force:",omitempty"`
	BillingGeocodeAccuracy  string   `force:",omitempty"`
	BillingLatitude         float64  `force:",omitempty"`
	BillingLongitude        float64  `force:",omitempty"`
	BillingPostalCode       string   `force:",omitempty"`
	BillingState            string   `force:",omitempty"`
	BillingStreet           string   `force:",omitempty"`
	CompanyAuthorizedById   string   `force:",omitempty"`
	CompanyAuthorizedDate   string   `force:",omitempty"`
	ContractId              string   `force:",omitempty"`
	CreatedById             string   `force:",omitempty"`
	CreatedDate             string   `force:",omitempty"`
	CustomerAuthorizedById  string   `force:",omitempty"`
	CustomerAuthorizedDate  string   `force:",omitempty"`
	Description             string   `force:",omitempty"`
	EffectiveDate           string   `force:",omitempty"`
	EndDate                 string   `force:",omitempty"`
	Id                      string   `force:",omitempty"`
	IsDeleted               bool     `force:",omitempty"`
	IsReductionOrder        bool     `force:",omitempty"`
	LastModifiedById        string   `force:",omitempty"`
	LastModifiedDate        string   `force:",omitempty"`
	LastReferencedDate      string   `force:",omitempty"`
	LastViewedDate          string   `force:",omitempty"`
	Name                    string   `force:",omitempty"`
	OrderNumber             string   `force:",omitempty"`
	OrderReferenceNumber    string   `force:",omitempty"`
	OriginalOrderId         string   `force:",omitempty"`
	OwnerId                 string   `force:",omitempty"`
	PoDate                  string   `force:",omitempty"`
	PoNumber                string   `force:",omitempty"`
	Pricebook2Id            string   `force:",omitempty"`
	ShipToContactId         string   `force:",omitempty"`
	ShippingAddress         *Address `force:",omitempty"`
	ShippingCity            string   `force:",omitempty"`
	ShippingCountry         string   `force:",omitempty"`
	ShippingGeocodeAccuracy string   `force:",omitempty"`
	ShippingLatitude        float64  `force:",omitempty"`
	ShippingLongitude       float64  `force:",omitempty"`
	ShippingPostalCode      string   `force:",omitempty"`
	ShippingState           string   `force:",omitempty"`
	ShippingStreet          string   `force:",omitempty"`
	Status                  string   `force:",omitempty"`
	StatusCode              string   `force:",omitempty"`
	SystemModstamp          string   `force:",omitempty"`
	TotalAmount             string   `force:",omitempty"`
	Type                    string   `force:",omitempty"`
}

func (t *Order) ApiName() string {
	return "Order"
}

func (t *Order) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Order #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccountId: %v\n", t.AccountId))
	builder.WriteString(fmt.Sprintf("\tActivatedById: %v\n", t.ActivatedById))
	builder.WriteString(fmt.Sprintf("\tActivatedDate: %v\n", t.ActivatedDate))
	builder.WriteString(fmt.Sprintf("\tBillToContactId: %v\n", t.BillToContactId))
	builder.WriteString(fmt.Sprintf("\tBillingAddress: %v\n", t.BillingAddress))
	builder.WriteString(fmt.Sprintf("\tBillingCity: %v\n", t.BillingCity))
	builder.WriteString(fmt.Sprintf("\tBillingCountry: %v\n", t.BillingCountry))
	builder.WriteString(fmt.Sprintf("\tBillingGeocodeAccuracy: %v\n", t.BillingGeocodeAccuracy))
	builder.WriteString(fmt.Sprintf("\tBillingLatitude: %v\n", t.BillingLatitude))
	builder.WriteString(fmt.Sprintf("\tBillingLongitude: %v\n", t.BillingLongitude))
	builder.WriteString(fmt.Sprintf("\tBillingPostalCode: %v\n", t.BillingPostalCode))
	builder.WriteString(fmt.Sprintf("\tBillingState: %v\n", t.BillingState))
	builder.WriteString(fmt.Sprintf("\tBillingStreet: %v\n", t.BillingStreet))
	builder.WriteString(fmt.Sprintf("\tCompanyAuthorizedById: %v\n", t.CompanyAuthorizedById))
	builder.WriteString(fmt.Sprintf("\tCompanyAuthorizedDate: %v\n", t.CompanyAuthorizedDate))
	builder.WriteString(fmt.Sprintf("\tContractId: %v\n", t.ContractId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCustomerAuthorizedById: %v\n", t.CustomerAuthorizedById))
	builder.WriteString(fmt.Sprintf("\tCustomerAuthorizedDate: %v\n", t.CustomerAuthorizedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tEffectiveDate: %v\n", t.EffectiveDate))
	builder.WriteString(fmt.Sprintf("\tEndDate: %v\n", t.EndDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsReductionOrder: %v\n", t.IsReductionOrder))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOrderNumber: %v\n", t.OrderNumber))
	builder.WriteString(fmt.Sprintf("\tOrderReferenceNumber: %v\n", t.OrderReferenceNumber))
	builder.WriteString(fmt.Sprintf("\tOriginalOrderId: %v\n", t.OriginalOrderId))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tPoDate: %v\n", t.PoDate))
	builder.WriteString(fmt.Sprintf("\tPoNumber: %v\n", t.PoNumber))
	builder.WriteString(fmt.Sprintf("\tPricebook2Id: %v\n", t.Pricebook2Id))
	builder.WriteString(fmt.Sprintf("\tShipToContactId: %v\n", t.ShipToContactId))
	builder.WriteString(fmt.Sprintf("\tShippingAddress: %v\n", t.ShippingAddress))
	builder.WriteString(fmt.Sprintf("\tShippingCity: %v\n", t.ShippingCity))
	builder.WriteString(fmt.Sprintf("\tShippingCountry: %v\n", t.ShippingCountry))
	builder.WriteString(fmt.Sprintf("\tShippingGeocodeAccuracy: %v\n", t.ShippingGeocodeAccuracy))
	builder.WriteString(fmt.Sprintf("\tShippingLatitude: %v\n", t.ShippingLatitude))
	builder.WriteString(fmt.Sprintf("\tShippingLongitude: %v\n", t.ShippingLongitude))
	builder.WriteString(fmt.Sprintf("\tShippingPostalCode: %v\n", t.ShippingPostalCode))
	builder.WriteString(fmt.Sprintf("\tShippingState: %v\n", t.ShippingState))
	builder.WriteString(fmt.Sprintf("\tShippingStreet: %v\n", t.ShippingStreet))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tStatusCode: %v\n", t.StatusCode))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTotalAmount: %v\n", t.TotalAmount))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type OrderQueryResponse struct {
	BaseQuery
	Records []Order `json:"Records" force:"records"`
}
