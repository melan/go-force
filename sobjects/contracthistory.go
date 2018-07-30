// This file was generated for SObject ContractHistory, API Version v43.0 at 2018-07-30 03:47:38.212989218 -0400 EDT m=+24.556517082

package sobjects

import (
	"fmt"
	"strings"
)

type ContractHistory struct {
	BaseSObject
	ContractId  string `force:",omitempty"`
	CreatedById string `force:",omitempty"`
	CreatedDate string `force:",omitempty"`
	Field       string `force:",omitempty"`
	Id          string `force:",omitempty"`
	IsDeleted   bool   `force:",omitempty"`
	NewValue    string `force:",omitempty"`
	OldValue    string `force:",omitempty"`
}

func (t *ContractHistory) ApiName() string {
	return "ContractHistory"
}

func (t *ContractHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContractHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContractId: %v\n", t.ContractId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))

	return builder.String()
}

type ContractHistoryQueryResponse struct {
	BaseQuery
	Records []ContractHistory `json:"Records" force:"records"`
}
