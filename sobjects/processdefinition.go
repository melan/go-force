// This file was generated for SObject ProcessDefinition, API Version v43.0 at 2018-07-30 03:47:31.578458672 -0400 EDT m=+17.921737582

package sobjects

import (
	"fmt"
	"strings"
)

type ProcessDefinition struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Description      string `force:",omitempty"`
	DeveloperName    string `force:",omitempty"`
	Id               string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	LockType         string `force:",omitempty"`
	Name             string `force:",omitempty"`
	State            string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	TableEnumOrId    string `force:",omitempty"`
	Type             string `force:",omitempty"`
}

func (t *ProcessDefinition) ApiName() string {
	return "ProcessDefinition"
}

func (t *ProcessDefinition) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ProcessDefinition #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLockType: %v\n", t.LockType))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tState: %v\n", t.State))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTableEnumOrId: %v\n", t.TableEnumOrId))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type ProcessDefinitionQueryResponse struct {
	BaseQuery
	Records []ProcessDefinition `json:"Records" force:"records"`
}
