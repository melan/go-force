// This file was generated for SObject IconDefinition, API Version v43.0 at 2018-07-30 03:47:41.831918126 -0400 EDT m=+28.175581787

package sobjects

import (
	"fmt"
	"strings"
)

type IconDefinition struct {
	BaseSObject
	ContentType     string `force:",omitempty"`
	DurableId       string `force:",omitempty"`
	Height          int    `force:",omitempty"`
	Id              string `force:",omitempty"`
	TabDefinitionId string `force:",omitempty"`
	Theme           string `force:",omitempty"`
	Url             string `force:",omitempty"`
	Width           int    `force:",omitempty"`
}

func (t *IconDefinition) ApiName() string {
	return "IconDefinition"
}

func (t *IconDefinition) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("IconDefinition #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentType: %v\n", t.ContentType))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tHeight: %v\n", t.Height))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tTabDefinitionId: %v\n", t.TabDefinitionId))
	builder.WriteString(fmt.Sprintf("\tTheme: %v\n", t.Theme))
	builder.WriteString(fmt.Sprintf("\tUrl: %v\n", t.Url))
	builder.WriteString(fmt.Sprintf("\tWidth: %v\n", t.Width))

	return builder.String()
}

type IconDefinitionQueryResponse struct {
	BaseQuery
	Records []IconDefinition `json:"Records" force:"records"`
}
