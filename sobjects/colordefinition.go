// This file was generated for SObject ColorDefinition, API Version v43.0 at 2018-07-30 03:47:36.46454492 -0400 EDT m=+22.808007176

package sobjects

import (
	"fmt"
	"strings"
)

type ColorDefinition struct {
	BaseSObject
	Color           string `force:",omitempty"`
	Context         string `force:",omitempty"`
	DurableId       string `force:",omitempty"`
	Id              string `force:",omitempty"`
	TabDefinitionId string `force:",omitempty"`
	Theme           string `force:",omitempty"`
}

func (t *ColorDefinition) ApiName() string {
	return "ColorDefinition"
}

func (t *ColorDefinition) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ColorDefinition #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tColor: %v\n", t.Color))
	builder.WriteString(fmt.Sprintf("\tContext: %v\n", t.Context))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tTabDefinitionId: %v\n", t.TabDefinitionId))
	builder.WriteString(fmt.Sprintf("\tTheme: %v\n", t.Theme))

	return builder.String()
}

type ColorDefinitionQueryResponse struct {
	BaseQuery
	Records []ColorDefinition `json:"Records" force:"records"`
}
