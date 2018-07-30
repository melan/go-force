// This file was generated for SObject ApexTestResult, API Version v43.0 at 2018-07-30 03:48:03.753424754 -0400 EDT m=+50.097910996

package sobjects

import (
	"fmt"
	"strings"
)

type ApexTestResult struct {
	BaseSObject
	ApexClassId         string `force:",omitempty"`
	ApexLogId           string `force:",omitempty"`
	ApexTestRunResultId string `force:",omitempty"`
	AsyncApexJobId      string `force:",omitempty"`
	Id                  string `force:",omitempty"`
	Message             string `force:",omitempty"`
	MethodName          string `force:",omitempty"`
	Outcome             string `force:",omitempty"`
	QueueItemId         string `force:",omitempty"`
	RunTime             int    `force:",omitempty"`
	StackTrace          string `force:",omitempty"`
	SystemModstamp      string `force:",omitempty"`
	TestTimestamp       string `force:",omitempty"`
}

func (t *ApexTestResult) ApiName() string {
	return "ApexTestResult"
}

func (t *ApexTestResult) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ApexTestResult #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApexClassId: %v\n", t.ApexClassId))
	builder.WriteString(fmt.Sprintf("\tApexLogId: %v\n", t.ApexLogId))
	builder.WriteString(fmt.Sprintf("\tApexTestRunResultId: %v\n", t.ApexTestRunResultId))
	builder.WriteString(fmt.Sprintf("\tAsyncApexJobId: %v\n", t.AsyncApexJobId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tMessage: %v\n", t.Message))
	builder.WriteString(fmt.Sprintf("\tMethodName: %v\n", t.MethodName))
	builder.WriteString(fmt.Sprintf("\tOutcome: %v\n", t.Outcome))
	builder.WriteString(fmt.Sprintf("\tQueueItemId: %v\n", t.QueueItemId))
	builder.WriteString(fmt.Sprintf("\tRunTime: %v\n", t.RunTime))
	builder.WriteString(fmt.Sprintf("\tStackTrace: %v\n", t.StackTrace))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTestTimestamp: %v\n", t.TestTimestamp))

	return builder.String()
}

type ApexTestResultQueryResponse struct {
	BaseQuery
	Records []ApexTestResult `json:"Records" force:"records"`
}
