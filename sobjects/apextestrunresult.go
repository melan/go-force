// This file was generated for SObject ApexTestRunResult, API Version v43.0 at 2018-07-30 03:48:06.346527851 -0400 EDT m=+52.691111397

package sobjects

import (
	"fmt"
	"strings"
)

type ApexTestRunResult struct {
	BaseSObject
	AsyncApexJobId   string `force:",omitempty"`
	ClassesCompleted int    `force:",omitempty"`
	ClassesEnqueued  int    `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	EndTime          string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsAllTests       bool   `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	JobName          string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	MethodsCompleted int    `force:",omitempty"`
	MethodsEnqueued  int    `force:",omitempty"`
	MethodsFailed    int    `force:",omitempty"`
	Source           string `force:",omitempty"`
	StartTime        string `force:",omitempty"`
	Status           string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	TestTime         int    `force:",omitempty"`
	UserId           string `force:",omitempty"`
}

func (t *ApexTestRunResult) ApiName() string {
	return "ApexTestRunResult"
}

func (t *ApexTestRunResult) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ApexTestRunResult #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAsyncApexJobId: %v\n", t.AsyncApexJobId))
	builder.WriteString(fmt.Sprintf("\tClassesCompleted: %v\n", t.ClassesCompleted))
	builder.WriteString(fmt.Sprintf("\tClassesEnqueued: %v\n", t.ClassesEnqueued))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tEndTime: %v\n", t.EndTime))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsAllTests: %v\n", t.IsAllTests))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tJobName: %v\n", t.JobName))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMethodsCompleted: %v\n", t.MethodsCompleted))
	builder.WriteString(fmt.Sprintf("\tMethodsEnqueued: %v\n", t.MethodsEnqueued))
	builder.WriteString(fmt.Sprintf("\tMethodsFailed: %v\n", t.MethodsFailed))
	builder.WriteString(fmt.Sprintf("\tSource: %v\n", t.Source))
	builder.WriteString(fmt.Sprintf("\tStartTime: %v\n", t.StartTime))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTestTime: %v\n", t.TestTime))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type ApexTestRunResultQueryResponse struct {
	BaseQuery
	Records []ApexTestRunResult `json:"Records" force:"records"`
}
