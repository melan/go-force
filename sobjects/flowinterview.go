// This file was generated for SObject FlowInterview, API Version v43.0 at 2018-07-30 03:47:19.391951852 -0400 EDT m=+5.734773475

package sobjects

import (
	"fmt"
	"strings"
)

type FlowInterview struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	CurrentElement   string `force:",omitempty"`
	Guid             string `force:",omitempty"`
	Id               string `force:",omitempty"`
	InterviewLabel   string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	OwnerId          string `force:",omitempty"`
	PauseLabel       string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *FlowInterview) ApiName() string {
	return "FlowInterview"
}

func (t *FlowInterview) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FlowInterview #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCurrentElement: %v\n", t.CurrentElement))
	builder.WriteString(fmt.Sprintf("\tGuid: %v\n", t.Guid))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInterviewLabel: %v\n", t.InterviewLabel))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tPauseLabel: %v\n", t.PauseLabel))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type FlowInterviewQueryResponse struct {
	BaseQuery
	Records []FlowInterview `json:"Records" force:"records"`
}
