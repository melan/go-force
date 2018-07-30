// This file was generated for SObject SearchPromotionRule, API Version v43.0 at 2018-07-30 03:47:24.929389613 -0400 EDT m=+11.272419024

package sobjects

import (
	"fmt"
	"strings"
)

type SearchPromotionRule struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Query            string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *SearchPromotionRule) ApiName() string {
	return "SearchPromotionRule"
}

func (t *SearchPromotionRule) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("SearchPromotionRule #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tQuery: %v\n", t.Query))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type SearchPromotionRuleQueryResponse struct {
	BaseQuery
	Records []SearchPromotionRule `json:"Records" force:"records"`
}
