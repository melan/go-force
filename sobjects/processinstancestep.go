// This file was generated for SObject ProcessInstanceStep, API Version v43.0 at 2018-07-30 03:48:07.384255419 -0400 EDT m=+53.728877904

package sobjects

import (
	"fmt"
	"strings"
)

type ProcessInstanceStep struct {
	BaseSObject
	ActorId              string  `force:",omitempty"`
	Comments             string  `force:",omitempty"`
	CreatedById          string  `force:",omitempty"`
	CreatedDate          string  `force:",omitempty"`
	ElapsedTimeInDays    float64 `force:",omitempty"`
	ElapsedTimeInHours   float64 `force:",omitempty"`
	ElapsedTimeInMinutes float64 `force:",omitempty"`
	Id                   string  `force:",omitempty"`
	OriginalActorId      string  `force:",omitempty"`
	ProcessInstanceId    string  `force:",omitempty"`
	StepNodeId           string  `force:",omitempty"`
	StepStatus           string  `force:",omitempty"`
	SystemModstamp       string  `force:",omitempty"`
}

func (t *ProcessInstanceStep) ApiName() string {
	return "ProcessInstanceStep"
}

func (t *ProcessInstanceStep) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ProcessInstanceStep #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tActorId: %v\n", t.ActorId))
	builder.WriteString(fmt.Sprintf("\tComments: %v\n", t.Comments))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tElapsedTimeInDays: %v\n", t.ElapsedTimeInDays))
	builder.WriteString(fmt.Sprintf("\tElapsedTimeInHours: %v\n", t.ElapsedTimeInHours))
	builder.WriteString(fmt.Sprintf("\tElapsedTimeInMinutes: %v\n", t.ElapsedTimeInMinutes))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tOriginalActorId: %v\n", t.OriginalActorId))
	builder.WriteString(fmt.Sprintf("\tProcessInstanceId: %v\n", t.ProcessInstanceId))
	builder.WriteString(fmt.Sprintf("\tStepNodeId: %v\n", t.StepNodeId))
	builder.WriteString(fmt.Sprintf("\tStepStatus: %v\n", t.StepStatus))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ProcessInstanceStepQueryResponse struct {
	BaseQuery
	Records []ProcessInstanceStep `json:"Records" force:"records"`
}
