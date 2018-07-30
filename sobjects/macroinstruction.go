// This file was generated for SObject MacroInstruction, API Version v43.0 at 2018-07-30 03:48:01.099118549 -0400 EDT m=+47.443505192

package sobjects

import (
	"fmt"
	"strings"
)

type MacroInstruction struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	MacroId          string `force:",omitempty"`
	Name             string `force:",omitempty"`
	Operation        string `force:",omitempty"`
	SortOrder        int    `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	Target           string `force:",omitempty"`
	Value            string `force:",omitempty"`
	ValueRecord      string `force:",omitempty"`
}

func (t *MacroInstruction) ApiName() string {
	return "MacroInstruction"
}

func (t *MacroInstruction) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("MacroInstruction #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMacroId: %v\n", t.MacroId))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOperation: %v\n", t.Operation))
	builder.WriteString(fmt.Sprintf("\tSortOrder: %v\n", t.SortOrder))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTarget: %v\n", t.Target))
	builder.WriteString(fmt.Sprintf("\tValue: %v\n", t.Value))
	builder.WriteString(fmt.Sprintf("\tValueRecord: %v\n", t.ValueRecord))

	return builder.String()
}

type MacroInstructionQueryResponse struct {
	BaseQuery
	Records []MacroInstruction `json:"Records" force:"records"`
}
