// This file was generated for SObject Product2, API Version v43.0 at 2018-07-30 03:47:35.745793123 -0400 EDT m=+22.089228408

package sobjects

import (
	"fmt"
	"strings"
)

type Product2 struct {
	BaseSObject
	CreatedById           string `force:",omitempty"`
	CreatedDate           string `force:",omitempty"`
	Description           string `force:",omitempty"`
	DisplayUrl            string `force:",omitempty"`
	ExternalDataSourceId  string `force:",omitempty"`
	ExternalId            string `force:",omitempty"`
	Family                string `force:",omitempty"`
	Id                    string `force:",omitempty"`
	IsActive              bool   `force:",omitempty"`
	IsArchived            bool   `force:",omitempty"`
	IsDeleted             bool   `force:",omitempty"`
	LastModifiedById      string `force:",omitempty"`
	LastModifiedDate      string `force:",omitempty"`
	LastReferencedDate    string `force:",omitempty"`
	LastViewedDate        string `force:",omitempty"`
	Name                  string `force:",omitempty"`
	ProductCode           string `force:",omitempty"`
	QuantityUnitOfMeasure string `force:",omitempty"`
	StockKeepingUnit      string `force:",omitempty"`
	SystemModstamp        string `force:",omitempty"`
}

func (t *Product2) ApiName() string {
	return "Product2"
}

func (t *Product2) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Product2 #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDisplayUrl: %v\n", t.DisplayUrl))
	builder.WriteString(fmt.Sprintf("\tExternalDataSourceId: %v\n", t.ExternalDataSourceId))
	builder.WriteString(fmt.Sprintf("\tExternalId: %v\n", t.ExternalId))
	builder.WriteString(fmt.Sprintf("\tFamily: %v\n", t.Family))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tIsArchived: %v\n", t.IsArchived))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tProductCode: %v\n", t.ProductCode))
	builder.WriteString(fmt.Sprintf("\tQuantityUnitOfMeasure: %v\n", t.QuantityUnitOfMeasure))
	builder.WriteString(fmt.Sprintf("\tStockKeepingUnit: %v\n", t.StockKeepingUnit))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type Product2QueryResponse struct {
	BaseQuery
	Records []Product2 `json:"Records" force:"records"`
}
