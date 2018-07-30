// This file was generated for SObject BusinessProcess, API Version v43.0 at 2018-07-30 03:47:39.685027236 -0400 EDT m=+26.028610337

package sobjects

import (
	"fmt"
	"strings"
)

type BusinessProcess struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Description      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsActive         bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	NamespacePrefix  string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	TableEnumOrId    string `force:",omitempty"`
}

func (t *BusinessProcess) ApiName() string {
	return "BusinessProcess"
}

func (t *BusinessProcess) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("BusinessProcess #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTableEnumOrId: %v\n", t.TableEnumOrId))

	return builder.String()
}

type BusinessProcessQueryResponse struct {
	BaseQuery
	Records []BusinessProcess `json:"Records" force:"records"`
}
