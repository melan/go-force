// This file was generated for SObject SearchLayout, API Version v43.0 at 2018-07-30 03:47:26.86311196 -0400 EDT m=+13.206213932

package sobjects

import (
	"fmt"
	"strings"
)

type SearchLayout struct {
	BaseSObject
	ButtonsDisplayed   string `force:",omitempty"`
	DurableId          string `force:",omitempty"`
	EntityDefinitionId string `force:",omitempty"`
	FieldsDisplayed    string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	Label              string `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	LayoutType         string `force:",omitempty"`
}

func (t *SearchLayout) ApiName() string {
	return "SearchLayout"
}

func (t *SearchLayout) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("SearchLayout #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tButtonsDisplayed: %v\n", t.ButtonsDisplayed))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tEntityDefinitionId: %v\n", t.EntityDefinitionId))
	builder.WriteString(fmt.Sprintf("\tFieldsDisplayed: %v\n", t.FieldsDisplayed))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLabel: %v\n", t.Label))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLayoutType: %v\n", t.LayoutType))

	return builder.String()
}

type SearchLayoutQueryResponse struct {
	BaseQuery
	Records []SearchLayout `json:"Records" force:"records"`
}
