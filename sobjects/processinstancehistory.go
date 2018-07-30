// This file was generated for SObject ProcessInstanceHistory, API Version v43.0 at 2018-07-30 03:47:22.162998645 -0400 EDT m=+8.505924250

package sobjects

import (
	"fmt"
	"strings"
)

type ProcessInstanceHistory struct {
	BaseSObject
	ActorId              string  `force:",omitempty"`
	Comments             string  `force:",omitempty"`
	CreatedById          string  `force:",omitempty"`
	CreatedDate          string  `force:",omitempty"`
	ElapsedTimeInDays    float64 `force:",omitempty"`
	ElapsedTimeInHours   float64 `force:",omitempty"`
	ElapsedTimeInMinutes float64 `force:",omitempty"`
	Id                   string  `force:",omitempty"`
	IsDeleted            bool    `force:",omitempty"`
	IsPending            bool    `force:",omitempty"`
	OriginalActorId      string  `force:",omitempty"`
	ProcessInstanceId    string  `force:",omitempty"`
	ProcessNodeId        string  `force:",omitempty"`
	RemindersSent        int     `force:",omitempty"`
	StepStatus           string  `force:",omitempty"`
	SystemModstamp       string  `force:",omitempty"`
	TargetObjectId       string  `force:",omitempty"`
}

func (t *ProcessInstanceHistory) ApiName() string {
	return "ProcessInstanceHistory"
}

func (t *ProcessInstanceHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ProcessInstanceHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tActorId: %v\n", t.ActorId))
	builder.WriteString(fmt.Sprintf("\tComments: %v\n", t.Comments))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tElapsedTimeInDays: %v\n", t.ElapsedTimeInDays))
	builder.WriteString(fmt.Sprintf("\tElapsedTimeInHours: %v\n", t.ElapsedTimeInHours))
	builder.WriteString(fmt.Sprintf("\tElapsedTimeInMinutes: %v\n", t.ElapsedTimeInMinutes))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsPending: %v\n", t.IsPending))
	builder.WriteString(fmt.Sprintf("\tOriginalActorId: %v\n", t.OriginalActorId))
	builder.WriteString(fmt.Sprintf("\tProcessInstanceId: %v\n", t.ProcessInstanceId))
	builder.WriteString(fmt.Sprintf("\tProcessNodeId: %v\n", t.ProcessNodeId))
	builder.WriteString(fmt.Sprintf("\tRemindersSent: %v\n", t.RemindersSent))
	builder.WriteString(fmt.Sprintf("\tStepStatus: %v\n", t.StepStatus))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTargetObjectId: %v\n", t.TargetObjectId))

	return builder.String()
}

type ProcessInstanceHistoryQueryResponse struct {
	BaseQuery
	Records []ProcessInstanceHistory `json:"Records" force:"records"`
}
