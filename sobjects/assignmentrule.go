// This file was generated for SObject AssignmentRule, API Version v43.0 at 2018-07-30 03:47:54.784073056 -0400 EDT m=+41.128222733

package sobjects

import (
	"fmt"
	"strings"
)

type AssignmentRule struct {
	BaseSObject
	Active           bool   `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	SobjectType      string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *AssignmentRule) ApiName() string {
	return "AssignmentRule"
}

func (t *AssignmentRule) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AssignmentRule #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tActive: %v\n", t.Active))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tSobjectType: %v\n", t.SobjectType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type AssignmentRuleQueryResponse struct {
	BaseQuery
	Records []AssignmentRule `json:"Records" force:"records"`
}
