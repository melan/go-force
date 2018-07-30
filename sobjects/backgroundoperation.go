// This file was generated for SObject BackgroundOperation, API Version v43.0 at 2018-07-30 03:47:17.502900191 -0400 EDT m=+3.845650929

package sobjects

import (
	"fmt"
	"strings"
)

type BackgroundOperation struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Error            string `force:",omitempty"`
	ExecutionGroup   string `force:",omitempty"`
	ExpiresAt        string `force:",omitempty"`
	FinishedAt       string `force:",omitempty"`
	GroupLeaderId    string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	NumFollowers     int    `force:",omitempty"`
	ParentKey        string `force:",omitempty"`
	ProcessAfter     string `force:",omitempty"`
	RetryBackoff     int    `force:",omitempty"`
	RetryCount       int    `force:",omitempty"`
	RetryLimit       int    `force:",omitempty"`
	SequenceGroup    string `force:",omitempty"`
	SequenceNumber   int    `force:",omitempty"`
	StartedAt        string `force:",omitempty"`
	Status           string `force:",omitempty"`
	SubmittedAt      string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	Timeout          int    `force:",omitempty"`
	WorkerUri        string `force:",omitempty"`
}

func (t *BackgroundOperation) ApiName() string {
	return "BackgroundOperation"
}

func (t *BackgroundOperation) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("BackgroundOperation #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tError: %v\n", t.Error))
	builder.WriteString(fmt.Sprintf("\tExecutionGroup: %v\n", t.ExecutionGroup))
	builder.WriteString(fmt.Sprintf("\tExpiresAt: %v\n", t.ExpiresAt))
	builder.WriteString(fmt.Sprintf("\tFinishedAt: %v\n", t.FinishedAt))
	builder.WriteString(fmt.Sprintf("\tGroupLeaderId: %v\n", t.GroupLeaderId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNumFollowers: %v\n", t.NumFollowers))
	builder.WriteString(fmt.Sprintf("\tParentKey: %v\n", t.ParentKey))
	builder.WriteString(fmt.Sprintf("\tProcessAfter: %v\n", t.ProcessAfter))
	builder.WriteString(fmt.Sprintf("\tRetryBackoff: %v\n", t.RetryBackoff))
	builder.WriteString(fmt.Sprintf("\tRetryCount: %v\n", t.RetryCount))
	builder.WriteString(fmt.Sprintf("\tRetryLimit: %v\n", t.RetryLimit))
	builder.WriteString(fmt.Sprintf("\tSequenceGroup: %v\n", t.SequenceGroup))
	builder.WriteString(fmt.Sprintf("\tSequenceNumber: %v\n", t.SequenceNumber))
	builder.WriteString(fmt.Sprintf("\tStartedAt: %v\n", t.StartedAt))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSubmittedAt: %v\n", t.SubmittedAt))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTimeout: %v\n", t.Timeout))
	builder.WriteString(fmt.Sprintf("\tWorkerUri: %v\n", t.WorkerUri))

	return builder.String()
}

type BackgroundOperationQueryResponse struct {
	BaseQuery
	Records []BackgroundOperation `json:"Records" force:"records"`
}
