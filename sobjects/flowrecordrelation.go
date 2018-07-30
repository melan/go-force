// This file was generated for SObject FlowRecordRelation, API Version v43.0 at 2018-07-30 03:48:08.473808081 -0400 EDT m=+54.818471451

package sobjects

import (
	"fmt"
	"strings"
)

type FlowRecordRelation struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	ParentId         string `force:",omitempty"`
	RelatedRecordId  string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *FlowRecordRelation) ApiName() string {
	return "FlowRecordRelation"
}

func (t *FlowRecordRelation) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FlowRecordRelation #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tRelatedRecordId: %v\n", t.RelatedRecordId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type FlowRecordRelationQueryResponse struct {
	BaseQuery
	Records []FlowRecordRelation `json:"Records" force:"records"`
}
