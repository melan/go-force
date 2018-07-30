// This file was generated for SObject ProcessInstance, API Version v43.0 at 2018-07-30 03:48:09.283708775 -0400 EDT m=+55.628402535

package sobjects

import (
	"fmt"
	"strings"
)

type ProcessInstance struct {
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
	ProcessDefinitionId  string  `force:",omitempty"`
	Status               string  `force:",omitempty"`
	SubmittedById        string  `force:",omitempty"`
	SystemModstamp       string  `force:",omitempty"`
	TargetObjectId       string  `force:",omitempty"`
}

func (t *ProcessInstance) ApiName() string {
	return "ProcessInstance"
}

func (t *ProcessInstance) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ProcessInstance #%s - %s\n", t.Id, t.Name))
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
	builder.WriteString(fmt.Sprintf("\tProcessDefinitionId: %v\n", t.ProcessDefinitionId))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSubmittedById: %v\n", t.SubmittedById))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTargetObjectId: %v\n", t.TargetObjectId))

	return builder.String()
}

type ProcessInstanceQueryResponse struct {
	BaseQuery
	Records []ProcessInstance `json:"Records" force:"records"`
}
