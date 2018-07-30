// This file was generated for SObject AccountHistory, API Version v43.0 at 2018-07-30 03:47:40.096680938 -0400 EDT m=+26.440279486

package sobjects

import (
	"fmt"
	"strings"
)

type AccountHistory struct {
	BaseSObject
	AccountId   string `force:",omitempty"`
	CreatedById string `force:",omitempty"`
	CreatedDate string `force:",omitempty"`
	Field       string `force:",omitempty"`
	Id          string `force:",omitempty"`
	IsDeleted   bool   `force:",omitempty"`
	NewValue    string `force:",omitempty"`
	OldValue    string `force:",omitempty"`
}

func (t *AccountHistory) ApiName() string {
	return "AccountHistory"
}

func (t *AccountHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AccountHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccountId: %v\n", t.AccountId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))

	return builder.String()
}

type AccountHistoryQueryResponse struct {
	BaseQuery
	Records []AccountHistory `json:"Records" force:"records"`
}
