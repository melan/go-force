// This file was generated for SObject MacroHistory, API Version v43.0 at 2018-07-30 03:47:22.023445909 -0400 EDT m=+8.366366277

package sobjects

import (
	"fmt"
	"strings"
)

type MacroHistory struct {
	BaseSObject
	CreatedById string `force:",omitempty"`
	CreatedDate string `force:",omitempty"`
	Field       string `force:",omitempty"`
	Id          string `force:",omitempty"`
	IsDeleted   bool   `force:",omitempty"`
	MacroId     string `force:",omitempty"`
	NewValue    string `force:",omitempty"`
	OldValue    string `force:",omitempty"`
}

func (t *MacroHistory) ApiName() string {
	return "MacroHistory"
}

func (t *MacroHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("MacroHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tMacroId: %v\n", t.MacroId))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))

	return builder.String()
}

type MacroHistoryQueryResponse struct {
	BaseQuery
	Records []MacroHistory `json:"Records" force:"records"`
}
