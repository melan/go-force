// This file was generated for SObject ContractStatus, API Version v43.0 at 2018-07-30 03:47:27.104336789 -0400 EDT m=+13.447447813

package sobjects

import (
	"fmt"
	"strings"
)

type ContractStatus struct {
	BaseSObject
	ApiName          string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDefault        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	MasterLabel      string `force:",omitempty"`
	SortOrder        int    `force:",omitempty"`
	StatusCode       string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *ContractStatus) ApiName() string {
	return "ContractStatus"
}

func (t *ContractStatus) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContractStatus #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApiName: %v\n", t.ApiName))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDefault: %v\n", t.IsDefault))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tSortOrder: %v\n", t.SortOrder))
	builder.WriteString(fmt.Sprintf("\tStatusCode: %v\n", t.StatusCode))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ContractStatusQueryResponse struct {
	BaseQuery
	Records []ContractStatus `json:"Records" force:"records"`
}
