// This file was generated for SObject AppTabMember, API Version v43.0 at 2018-07-30 03:47:27.836491455 -0400 EDT m=+14.179629952

package sobjects

import (
	"fmt"
	"strings"
)

type AppTabMember struct {
	BaseSObject
	AppDefinitionId      string `force:",omitempty"`
	DurableId            string `force:",omitempty"`
	Id                   string `force:",omitempty"`
	SortOrder            int    `force:",omitempty"`
	TabDefinitionId      string `force:",omitempty"`
	WorkspaceDriverField string `force:",omitempty"`
}

func (t *AppTabMember) ApiName() string {
	return "AppTabMember"
}

func (t *AppTabMember) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AppTabMember #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAppDefinitionId: %v\n", t.AppDefinitionId))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tSortOrder: %v\n", t.SortOrder))
	builder.WriteString(fmt.Sprintf("\tTabDefinitionId: %v\n", t.TabDefinitionId))
	builder.WriteString(fmt.Sprintf("\tWorkspaceDriverField: %v\n", t.WorkspaceDriverField))

	return builder.String()
}

type AppTabMemberQueryResponse struct {
	BaseQuery
	Records []AppTabMember `json:"Records" force:"records"`
}
