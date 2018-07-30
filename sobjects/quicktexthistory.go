// This file was generated for SObject QuickTextHistory, API Version v43.0 at 2018-07-30 03:47:21.120883878 -0400 EDT m=+7.463770378

package sobjects

import (
	"fmt"
	"strings"
)

type QuickTextHistory struct {
	BaseSObject
	CreatedById string `force:",omitempty"`
	CreatedDate string `force:",omitempty"`
	Field       string `force:",omitempty"`
	Id          string `force:",omitempty"`
	IsDeleted   bool   `force:",omitempty"`
	NewValue    string `force:",omitempty"`
	OldValue    string `force:",omitempty"`
	QuickTextId string `force:",omitempty"`
}

func (t *QuickTextHistory) ApiName() string {
	return "QuickTextHistory"
}

func (t *QuickTextHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("QuickTextHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))
	builder.WriteString(fmt.Sprintf("\tQuickTextId: %v\n", t.QuickTextId))

	return builder.String()
}

type QuickTextHistoryQueryResponse struct {
	BaseQuery
	Records []QuickTextHistory `json:"Records" force:"records"`
}
