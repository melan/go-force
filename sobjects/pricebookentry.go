// This file was generated for SObject PricebookEntry, API Version v43.0 at 2018-07-30 03:47:44.756655435 -0400 EDT m=+31.100428844

package sobjects

import (
	"fmt"
	"strings"
)

type PricebookEntry struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsActive         bool   `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	Pricebook2Id     string `force:",omitempty"`
	Product2Id       string `force:",omitempty"`
	ProductCode      string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	UnitPrice        string `force:",omitempty"`
	UseStandardPrice bool   `force:",omitempty"`
}

func (t *PricebookEntry) ApiName() string {
	return "PricebookEntry"
}

func (t *PricebookEntry) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("PricebookEntry #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tPricebook2Id: %v\n", t.Pricebook2Id))
	builder.WriteString(fmt.Sprintf("\tProduct2Id: %v\n", t.Product2Id))
	builder.WriteString(fmt.Sprintf("\tProductCode: %v\n", t.ProductCode))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUnitPrice: %v\n", t.UnitPrice))
	builder.WriteString(fmt.Sprintf("\tUseStandardPrice: %v\n", t.UseStandardPrice))

	return builder.String()
}

type PricebookEntryQueryResponse struct {
	BaseQuery
	Records []PricebookEntry `json:"Records" force:"records"`
}
