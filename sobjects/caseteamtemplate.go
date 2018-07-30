// This file was generated for SObject CaseTeamTemplate, API Version v43.0 at 2018-07-30 03:47:39.961302704 -0400 EDT m=+26.304896172

package sobjects

import (
	"fmt"
	"strings"
)

type CaseTeamTemplate struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Description      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *CaseTeamTemplate) ApiName() string {
	return "CaseTeamTemplate"
}

func (t *CaseTeamTemplate) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CaseTeamTemplate #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CaseTeamTemplateQueryResponse struct {
	BaseQuery
	Records []CaseTeamTemplate `json:"Records" force:"records"`
}
