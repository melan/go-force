// This file was generated for SObject OwnerChangeOptionInfo, API Version v43.0 at 2018-07-30 03:48:10.54309692 -0400 EDT m=+56.887837938

package sobjects

import (
	"fmt"
	"strings"
)

type OwnerChangeOptionInfo struct {
	BaseSObject
	DefaultValue       bool   `force:",omitempty"`
	DurableId          string `force:",omitempty"`
	EntityDefinitionId string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsEditable         bool   `force:",omitempty"`
	Label              string `force:",omitempty"`
	Name               string `force:",omitempty"`
}

func (t *OwnerChangeOptionInfo) ApiName() string {
	return "OwnerChangeOptionInfo"
}

func (t *OwnerChangeOptionInfo) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OwnerChangeOptionInfo #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tDefaultValue: %v\n", t.DefaultValue))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tEntityDefinitionId: %v\n", t.EntityDefinitionId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsEditable: %v\n", t.IsEditable))
	builder.WriteString(fmt.Sprintf("\tLabel: %v\n", t.Label))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))

	return builder.String()
}

type OwnerChangeOptionInfoQueryResponse struct {
	BaseQuery
	Records []OwnerChangeOptionInfo `json:"Records" force:"records"`
}
