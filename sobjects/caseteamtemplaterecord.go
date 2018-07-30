// This file was generated for SObject CaseTeamTemplateRecord, API Version v43.0 at 2018-07-30 03:47:36.232407394 -0400 EDT m=+22.575860939

package sobjects

import (
	"fmt"
	"strings"
)

type CaseTeamTemplateRecord struct {
	BaseSObject
	CreatedById    string `force:",omitempty"`
	CreatedDate    string `force:",omitempty"`
	Id             string `force:",omitempty"`
	ParentId       string `force:",omitempty"`
	SystemModstamp string `force:",omitempty"`
	TeamTemplateId string `force:",omitempty"`
}

func (t *CaseTeamTemplateRecord) ApiName() string {
	return "CaseTeamTemplateRecord"
}

func (t *CaseTeamTemplateRecord) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CaseTeamTemplateRecord #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTeamTemplateId: %v\n", t.TeamTemplateId))

	return builder.String()
}

type CaseTeamTemplateRecordQueryResponse struct {
	BaseQuery
	Records []CaseTeamTemplateRecord `json:"Records" force:"records"`
}
