// This file was generated for SObject ProcessNode, API Version v43.0 at 2018-07-30 03:47:53.448334724 -0400 EDT m=+39.792434278

package sobjects

import (
	"fmt"
	"strings"
)

type ProcessNode struct {
	BaseSObject
	Description         string `force:",omitempty"`
	DeveloperName       string `force:",omitempty"`
	Id                  string `force:",omitempty"`
	Name                string `force:",omitempty"`
	ProcessDefinitionId string `force:",omitempty"`
	SystemModstamp      string `force:",omitempty"`
}

func (t *ProcessNode) ApiName() string {
	return "ProcessNode"
}

func (t *ProcessNode) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ProcessNode #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tProcessDefinitionId: %v\n", t.ProcessDefinitionId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ProcessNodeQueryResponse struct {
	BaseQuery
	Records []ProcessNode `json:"Records" force:"records"`
}
