// This file was generated for SObject FlowStageRelation, API Version v43.0 at 2018-07-30 03:47:30.385921669 -0400 EDT m=+16.729155831

package sobjects

import (
	"fmt"
	"strings"
)

type FlowStageRelation struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	FlexIndex        string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	ParentId         string `force:",omitempty"`
	StageLabel       string `force:",omitempty"`
	StageOrder       int    `force:",omitempty"`
	StageType        string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *FlowStageRelation) ApiName() string {
	return "FlowStageRelation"
}

func (t *FlowStageRelation) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FlowStageRelation #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFlexIndex: %v\n", t.FlexIndex))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tStageLabel: %v\n", t.StageLabel))
	builder.WriteString(fmt.Sprintf("\tStageOrder: %v\n", t.StageOrder))
	builder.WriteString(fmt.Sprintf("\tStageType: %v\n", t.StageType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type FlowStageRelationQueryResponse struct {
	BaseQuery
	Records []FlowStageRelation `json:"Records" force:"records"`
}
