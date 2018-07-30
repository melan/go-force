// This file was generated for SObject ApexTestQueueItem, API Version v43.0 at 2018-07-30 03:47:19.674375195 -0400 EDT m=+6.017207416

package sobjects

import (
	"fmt"
	"strings"
)

type ApexTestQueueItem struct {
	BaseSObject
	ApexClassId            string `force:",omitempty"`
	CreatedById            string `force:",omitempty"`
	CreatedDate            string `force:",omitempty"`
	ExtendedStatus         string `force:",omitempty"`
	Id                     string `force:",omitempty"`
	ParentJobId            string `force:",omitempty"`
	ShouldSkipCodeCoverage bool   `force:",omitempty"`
	Status                 string `force:",omitempty"`
	SystemModstamp         string `force:",omitempty"`
	TestRunResultId        string `force:",omitempty"`
}

func (t *ApexTestQueueItem) ApiName() string {
	return "ApexTestQueueItem"
}

func (t *ApexTestQueueItem) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ApexTestQueueItem #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApexClassId: %v\n", t.ApexClassId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tExtendedStatus: %v\n", t.ExtendedStatus))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tParentJobId: %v\n", t.ParentJobId))
	builder.WriteString(fmt.Sprintf("\tShouldSkipCodeCoverage: %v\n", t.ShouldSkipCodeCoverage))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTestRunResultId: %v\n", t.TestRunResultId))

	return builder.String()
}

type ApexTestQueueItemQueryResponse struct {
	BaseQuery
	Records []ApexTestQueueItem `json:"Records" force:"records"`
}
