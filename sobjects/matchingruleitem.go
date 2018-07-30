// This file was generated for SObject MatchingRuleItem, API Version v43.0 at 2018-07-30 03:47:56.508272345 -0400 EDT m=+42.852486720

package sobjects

import (
	"fmt"
	"strings"
)

type MatchingRuleItem struct {
	BaseSObject
	BlankValueBehavior string `force:",omitempty"`
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	Field              string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	MatchingMethod     string `force:",omitempty"`
	MatchingRuleId     string `force:",omitempty"`
	SortOrder          int    `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
}

func (t *MatchingRuleItem) ApiName() string {
	return "MatchingRuleItem"
}

func (t *MatchingRuleItem) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("MatchingRuleItem #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tBlankValueBehavior: %v\n", t.BlankValueBehavior))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMatchingMethod: %v\n", t.MatchingMethod))
	builder.WriteString(fmt.Sprintf("\tMatchingRuleId: %v\n", t.MatchingRuleId))
	builder.WriteString(fmt.Sprintf("\tSortOrder: %v\n", t.SortOrder))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type MatchingRuleItemQueryResponse struct {
	BaseQuery
	Records []MatchingRuleItem `json:"Records" force:"records"`
}
