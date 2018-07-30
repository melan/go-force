// This file was generated for SObject OrderShare, API Version v43.0 at 2018-07-30 03:47:33.505774007 -0400 EDT m=+19.849125238

package sobjects

import (
	"fmt"
	"strings"
)

type OrderShare struct {
	BaseSObject
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	OrderAccessLevel string `force:",omitempty"`
	OrderId          string `force:",omitempty"`
	RowCause         string `force:",omitempty"`
	UserOrGroupId    string `force:",omitempty"`
}

func (t *OrderShare) ApiName() string {
	return "OrderShare"
}

func (t *OrderShare) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OrderShare #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tOrderAccessLevel: %v\n", t.OrderAccessLevel))
	builder.WriteString(fmt.Sprintf("\tOrderId: %v\n", t.OrderId))
	builder.WriteString(fmt.Sprintf("\tRowCause: %v\n", t.RowCause))
	builder.WriteString(fmt.Sprintf("\tUserOrGroupId: %v\n", t.UserOrGroupId))

	return builder.String()
}

type OrderShareQueryResponse struct {
	BaseQuery
	Records []OrderShare `json:"Records" force:"records"`
}
