// This file was generated for SObject TabDefinition, API Version v43.0 at 2018-07-30 03:47:40.46448863 -0400 EDT m=+26.808100979

package sobjects

import (
	"fmt"
	"strings"
)

type TabDefinition struct {
	BaseSObject
	DurableId              string `force:",omitempty"`
	Id                     string `force:",omitempty"`
	IsAvailableInAloha     bool   `force:",omitempty"`
	IsAvailableInLightning bool   `force:",omitempty"`
	IsCustom               bool   `force:",omitempty"`
	Label                  string `force:",omitempty"`
	Name                   string `force:",omitempty"`
	SobjectName            string `force:",omitempty"`
	Url                    string `force:",omitempty"`
}

func (t *TabDefinition) ApiName() string {
	return "TabDefinition"
}

func (t *TabDefinition) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("TabDefinition #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsAvailableInAloha: %v\n", t.IsAvailableInAloha))
	builder.WriteString(fmt.Sprintf("\tIsAvailableInLightning: %v\n", t.IsAvailableInLightning))
	builder.WriteString(fmt.Sprintf("\tIsCustom: %v\n", t.IsCustom))
	builder.WriteString(fmt.Sprintf("\tLabel: %v\n", t.Label))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tSobjectName: %v\n", t.SobjectName))
	builder.WriteString(fmt.Sprintf("\tUrl: %v\n", t.Url))

	return builder.String()
}

type TabDefinitionQueryResponse struct {
	BaseQuery
	Records []TabDefinition `json:"Records" force:"records"`
}
