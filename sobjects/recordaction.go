// This file was generated for SObject RecordAction, API Version v43.0 at 2018-07-30 03:48:11.726860535 -0400 EDT m=+58.071645972

package sobjects

import (
	"fmt"
	"strings"
)

type RecordAction struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	FlowDefinition   string `force:",omitempty"`
	FlowInterviewId  string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Order            int    `force:",omitempty"`
	Pinned           string `force:",omitempty"`
	RecordId         string `force:",omitempty"`
	Status           string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *RecordAction) ApiName() string {
	return "RecordAction"
}

func (t *RecordAction) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("RecordAction #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFlowDefinition: %v\n", t.FlowDefinition))
	builder.WriteString(fmt.Sprintf("\tFlowInterviewId: %v\n", t.FlowInterviewId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tOrder: %v\n", t.Order))
	builder.WriteString(fmt.Sprintf("\tPinned: %v\n", t.Pinned))
	builder.WriteString(fmt.Sprintf("\tRecordId: %v\n", t.RecordId))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type RecordActionQueryResponse struct {
	BaseQuery
	Records []RecordAction `json:"Records" force:"records"`
}
