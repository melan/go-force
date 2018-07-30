// This file was generated for SObject LeadStatus, API Version v43.0 at 2018-07-30 03:48:06.46352297 -0400 EDT m=+52.808110906

package sobjects

import (
	"fmt"
	"strings"
)

type LeadStatus struct {
	BaseSObject
	ApiName          string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsConverted      bool   `force:",omitempty"`
	IsDefault        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	MasterLabel      string `force:",omitempty"`
	SortOrder        int    `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *LeadStatus) ApiName() string {
	return "LeadStatus"
}

func (t *LeadStatus) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("LeadStatus #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApiName: %v\n", t.ApiName))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsConverted: %v\n", t.IsConverted))
	builder.WriteString(fmt.Sprintf("\tIsDefault: %v\n", t.IsDefault))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tSortOrder: %v\n", t.SortOrder))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type LeadStatusQueryResponse struct {
	BaseQuery
	Records []LeadStatus `json:"Records" force:"records"`
}
