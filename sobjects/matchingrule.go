// This file was generated for SObject MatchingRule, API Version v43.0 at 2018-07-30 03:47:43.638547621 -0400 EDT m=+29.982279074

package sobjects

import (
	"fmt"
	"strings"
)

type MatchingRule struct {
	BaseSObject
	BooleanFilter    string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Description      string `force:",omitempty"`
	DeveloperName    string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	Language         string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	MasterLabel      string `force:",omitempty"`
	MatchEngine      string `force:",omitempty"`
	NamespacePrefix  string `force:",omitempty"`
	RuleStatus       string `force:",omitempty"`
	SobjectSubtype   string `force:",omitempty"`
	SobjectType      string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *MatchingRule) ApiName() string {
	return "MatchingRule"
}

func (t *MatchingRule) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("MatchingRule #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tBooleanFilter: %v\n", t.BooleanFilter))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tMatchEngine: %v\n", t.MatchEngine))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tRuleStatus: %v\n", t.RuleStatus))
	builder.WriteString(fmt.Sprintf("\tSobjectSubtype: %v\n", t.SobjectSubtype))
	builder.WriteString(fmt.Sprintf("\tSobjectType: %v\n", t.SobjectType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type MatchingRuleQueryResponse struct {
	BaseQuery
	Records []MatchingRule `json:"Records" force:"records"`
}
