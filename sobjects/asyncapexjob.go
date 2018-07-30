// This file was generated for SObject AsyncApexJob, API Version v43.0 at 2018-07-30 03:48:07.797174821 -0400 EDT m=+54.141812801

package sobjects

import (
	"fmt"
	"strings"
)

type AsyncApexJob struct {
	BaseSObject
	ApexClassId         string `force:",omitempty"`
	CompletedDate       string `force:",omitempty"`
	CreatedById         string `force:",omitempty"`
	CreatedDate         string `force:",omitempty"`
	ExtendedStatus      string `force:",omitempty"`
	Id                  string `force:",omitempty"`
	JobItemsProcessed   int    `force:",omitempty"`
	JobType             string `force:",omitempty"`
	LastProcessed       string `force:",omitempty"`
	LastProcessedOffset int    `force:",omitempty"`
	MethodName          string `force:",omitempty"`
	NumberOfErrors      int    `force:",omitempty"`
	ParentJobId         string `force:",omitempty"`
	Status              string `force:",omitempty"`
	TotalJobItems       int    `force:",omitempty"`
}

func (t *AsyncApexJob) ApiName() string {
	return "AsyncApexJob"
}

func (t *AsyncApexJob) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AsyncApexJob #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApexClassId: %v\n", t.ApexClassId))
	builder.WriteString(fmt.Sprintf("\tCompletedDate: %v\n", t.CompletedDate))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tExtendedStatus: %v\n", t.ExtendedStatus))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tJobItemsProcessed: %v\n", t.JobItemsProcessed))
	builder.WriteString(fmt.Sprintf("\tJobType: %v\n", t.JobType))
	builder.WriteString(fmt.Sprintf("\tLastProcessed: %v\n", t.LastProcessed))
	builder.WriteString(fmt.Sprintf("\tLastProcessedOffset: %v\n", t.LastProcessedOffset))
	builder.WriteString(fmt.Sprintf("\tMethodName: %v\n", t.MethodName))
	builder.WriteString(fmt.Sprintf("\tNumberOfErrors: %v\n", t.NumberOfErrors))
	builder.WriteString(fmt.Sprintf("\tParentJobId: %v\n", t.ParentJobId))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tTotalJobItems: %v\n", t.TotalJobItems))

	return builder.String()
}

type AsyncApexJobQueryResponse struct {
	BaseQuery
	Records []AsyncApexJob `json:"Records" force:"records"`
}
