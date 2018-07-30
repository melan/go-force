// This file was generated for SObject OrderItem, API Version v43.0 at 2018-07-30 03:47:57.572955407 -0400 EDT m=+43.917209734

package sobjects

import (
	"fmt"
	"strings"
)

type OrderItem struct {
	BaseSObject
	AvailableQuantity   float64 `force:",omitempty"`
	CreatedById         string  `force:",omitempty"`
	CreatedDate         string  `force:",omitempty"`
	Description         string  `force:",omitempty"`
	EndDate             string  `force:",omitempty"`
	Id                  string  `force:",omitempty"`
	IsDeleted           bool    `force:",omitempty"`
	LastModifiedById    string  `force:",omitempty"`
	LastModifiedDate    string  `force:",omitempty"`
	ListPrice           string  `force:",omitempty"`
	OrderId             string  `force:",omitempty"`
	OrderItemNumber     string  `force:",omitempty"`
	OriginalOrderItemId string  `force:",omitempty"`
	PricebookEntryId    string  `force:",omitempty"`
	Product2Id          string  `force:",omitempty"`
	Quantity            float64 `force:",omitempty"`
	ServiceDate         string  `force:",omitempty"`
	SystemModstamp      string  `force:",omitempty"`
	TotalPrice          string  `force:",omitempty"`
	UnitPrice           string  `force:",omitempty"`
}

func (t *OrderItem) ApiName() string {
	return "OrderItem"
}

func (t *OrderItem) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OrderItem #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAvailableQuantity: %v\n", t.AvailableQuantity))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tEndDate: %v\n", t.EndDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tListPrice: %v\n", t.ListPrice))
	builder.WriteString(fmt.Sprintf("\tOrderId: %v\n", t.OrderId))
	builder.WriteString(fmt.Sprintf("\tOrderItemNumber: %v\n", t.OrderItemNumber))
	builder.WriteString(fmt.Sprintf("\tOriginalOrderItemId: %v\n", t.OriginalOrderItemId))
	builder.WriteString(fmt.Sprintf("\tPricebookEntryId: %v\n", t.PricebookEntryId))
	builder.WriteString(fmt.Sprintf("\tProduct2Id: %v\n", t.Product2Id))
	builder.WriteString(fmt.Sprintf("\tQuantity: %v\n", t.Quantity))
	builder.WriteString(fmt.Sprintf("\tServiceDate: %v\n", t.ServiceDate))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTotalPrice: %v\n", t.TotalPrice))
	builder.WriteString(fmt.Sprintf("\tUnitPrice: %v\n", t.UnitPrice))

	return builder.String()
}

type OrderItemQueryResponse struct {
	BaseQuery
	Records []OrderItem `json:"Records" force:"records"`
}
