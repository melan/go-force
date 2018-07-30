// This file was generated for SObject Pricebook2History, API Version v43.0 at 2018-07-30 03:48:04.002373816 -0400 EDT m=+50.346869400

package sobjects

import (
	"fmt"
	"strings"
)

type Pricebook2History struct {
	BaseSObject
	CreatedById  string `force:",omitempty"`
	CreatedDate  string `force:",omitempty"`
	Field        string `force:",omitempty"`
	Id           string `force:",omitempty"`
	IsDeleted    bool   `force:",omitempty"`
	NewValue     string `force:",omitempty"`
	OldValue     string `force:",omitempty"`
	Pricebook2Id string `force:",omitempty"`
}

func (t *Pricebook2History) ApiName() string {
	return "Pricebook2History"
}

func (t *Pricebook2History) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Pricebook2History #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))
	builder.WriteString(fmt.Sprintf("\tPricebook2Id: %v\n", t.Pricebook2Id))

	return builder.String()
}

type Pricebook2HistoryQueryResponse struct {
	BaseQuery
	Records []Pricebook2History `json:"Records" force:"records"`
}
