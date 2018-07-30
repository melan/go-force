// This file was generated for SObject Product2History, API Version v43.0 at 2018-07-30 03:48:08.176029262 -0400 EDT m=+54.520681458

package sobjects

import (
	"fmt"
	"strings"
)

type Product2History struct {
	BaseSObject
	CreatedById string `force:",omitempty"`
	CreatedDate string `force:",omitempty"`
	Field       string `force:",omitempty"`
	Id          string `force:",omitempty"`
	IsDeleted   bool   `force:",omitempty"`
	NewValue    string `force:",omitempty"`
	OldValue    string `force:",omitempty"`
	Product2Id  string `force:",omitempty"`
}

func (t *Product2History) ApiName() string {
	return "Product2History"
}

func (t *Product2History) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Product2History #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))
	builder.WriteString(fmt.Sprintf("\tProduct2Id: %v\n", t.Product2Id))

	return builder.String()
}

type Product2HistoryQueryResponse struct {
	BaseQuery
	Records []Product2History `json:"Records" force:"records"`
}
