// This file was generated for SObject ProcessInstanceNode, API Version v43.0 at 2018-07-30 03:47:17.781455759 -0400 EDT m=+4.124216949

package sobjects

import (
	"fmt"
	"strings"
)

type ProcessInstanceNode struct {
	BaseSObject
	CompletedDate        string  `force:",omitempty"`
	CreatedById          string  `force:",omitempty"`
	CreatedDate          string  `force:",omitempty"`
	ElapsedTimeInDays    float64 `force:",omitempty"`
	ElapsedTimeInHours   float64 `force:",omitempty"`
	ElapsedTimeInMinutes float64 `force:",omitempty"`
	Id                   string  `force:",omitempty"`
	IsDeleted            bool    `force:",omitempty"`
	LastActorId          string  `force:",omitempty"`
	LastModifiedById     string  `force:",omitempty"`
	LastModifiedDate     string  `force:",omitempty"`
	NodeStatus           string  `force:",omitempty"`
	ProcessInstanceId    string  `force:",omitempty"`
	ProcessNodeId        string  `force:",omitempty"`
	ProcessNodeName      string  `force:",omitempty"`
	SystemModstamp       string  `force:",omitempty"`
}

func (t *ProcessInstanceNode) ApiName() string {
	return "ProcessInstanceNode"
}

func (t *ProcessInstanceNode) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ProcessInstanceNode #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCompletedDate: %v\n", t.CompletedDate))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tElapsedTimeInDays: %v\n", t.ElapsedTimeInDays))
	builder.WriteString(fmt.Sprintf("\tElapsedTimeInHours: %v\n", t.ElapsedTimeInHours))
	builder.WriteString(fmt.Sprintf("\tElapsedTimeInMinutes: %v\n", t.ElapsedTimeInMinutes))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastActorId: %v\n", t.LastActorId))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tNodeStatus: %v\n", t.NodeStatus))
	builder.WriteString(fmt.Sprintf("\tProcessInstanceId: %v\n", t.ProcessInstanceId))
	builder.WriteString(fmt.Sprintf("\tProcessNodeId: %v\n", t.ProcessNodeId))
	builder.WriteString(fmt.Sprintf("\tProcessNodeName: %v\n", t.ProcessNodeName))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ProcessInstanceNodeQueryResponse struct {
	BaseQuery
	Records []ProcessInstanceNode `json:"Records" force:"records"`
}
