// This file was generated for SObject OrderHistory, API Version v43.0 at 2018-07-30 03:48:08.054050014 -0400 EDT m=+54.398697633

package sobjects

import (
	"fmt"
	"strings"
)

type OrderHistory struct {
	BaseSObject
	CreatedById string `force:",omitempty"`
	CreatedDate string `force:",omitempty"`
	Field       string `force:",omitempty"`
	Id          string `force:",omitempty"`
	IsDeleted   bool   `force:",omitempty"`
	NewValue    string `force:",omitempty"`
	OldValue    string `force:",omitempty"`
	OrderId     string `force:",omitempty"`
}

func (t *OrderHistory) ApiName() string {
	return "OrderHistory"
}

func (t *OrderHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OrderHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))
	builder.WriteString(fmt.Sprintf("\tOrderId: %v\n", t.OrderId))

	return builder.String()
}

type OrderHistoryQueryResponse struct {
	BaseQuery
	Records []OrderHistory `json:"Records" force:"records"`
}
