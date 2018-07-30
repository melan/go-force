// This file was generated for SObject ProcessInstanceWorkitem, API Version v43.0 at 2018-07-30 03:48:10.792652135 -0400 EDT m=+57.137402517

package sobjects

import (
	"fmt"
	"strings"
)

type ProcessInstanceWorkitem struct {
	BaseSObject
	ActorId              string  `force:",omitempty"`
	CreatedById          string  `force:",omitempty"`
	CreatedDate          string  `force:",omitempty"`
	ElapsedTimeInDays    float64 `force:",omitempty"`
	ElapsedTimeInHours   float64 `force:",omitempty"`
	ElapsedTimeInMinutes float64 `force:",omitempty"`
	Id                   string  `force:",omitempty"`
	IsDeleted            bool    `force:",omitempty"`
	OriginalActorId      string  `force:",omitempty"`
	ProcessInstanceId    string  `force:",omitempty"`
	SystemModstamp       string  `force:",omitempty"`
}

func (t *ProcessInstanceWorkitem) ApiName() string {
	return "ProcessInstanceWorkitem"
}

func (t *ProcessInstanceWorkitem) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ProcessInstanceWorkitem #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tActorId: %v\n", t.ActorId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tElapsedTimeInDays: %v\n", t.ElapsedTimeInDays))
	builder.WriteString(fmt.Sprintf("\tElapsedTimeInHours: %v\n", t.ElapsedTimeInHours))
	builder.WriteString(fmt.Sprintf("\tElapsedTimeInMinutes: %v\n", t.ElapsedTimeInMinutes))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tOriginalActorId: %v\n", t.OriginalActorId))
	builder.WriteString(fmt.Sprintf("\tProcessInstanceId: %v\n", t.ProcessInstanceId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ProcessInstanceWorkitemQueryResponse struct {
	BaseQuery
	Records []ProcessInstanceWorkitem `json:"Records" force:"records"`
}
