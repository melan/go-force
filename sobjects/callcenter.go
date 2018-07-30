// This file was generated for SObject CallCenter, API Version v43.0 at 2018-07-30 03:47:20.731256643 -0400 EDT m=+7.074128523

package sobjects

import (
	"fmt"
	"strings"
)

type CallCenter struct {
	BaseSObject
	AdapterUrl       string  `force:",omitempty"`
	CreatedById      string  `force:",omitempty"`
	CreatedDate      string  `force:",omitempty"`
	CustomSettings   string  `force:",omitempty"`
	Id               string  `force:",omitempty"`
	InternalName     string  `force:",omitempty"`
	LastModifiedById string  `force:",omitempty"`
	LastModifiedDate string  `force:",omitempty"`
	Name             string  `force:",omitempty"`
	SystemModstamp   string  `force:",omitempty"`
	Version          float64 `force:",omitempty"`
}

func (t *CallCenter) ApiName() string {
	return "CallCenter"
}

func (t *CallCenter) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CallCenter #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAdapterUrl: %v\n", t.AdapterUrl))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCustomSettings: %v\n", t.CustomSettings))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInternalName: %v\n", t.InternalName))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tVersion: %v\n", t.Version))

	return builder.String()
}

type CallCenterQueryResponse struct {
	BaseQuery
	Records []CallCenter `json:"Records" force:"records"`
}
