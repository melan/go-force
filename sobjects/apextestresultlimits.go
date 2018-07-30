// This file was generated for SObject ApexTestResultLimits, API Version v43.0 at 2018-07-30 03:48:04.968588847 -0400 EDT m=+51.313120687

package sobjects

import (
	"fmt"
	"strings"
)

type ApexTestResultLimits struct {
	BaseSObject
	ApexTestResultId string `force:",omitempty"`
	AsyncCalls       int    `force:",omitempty"`
	Callouts         int    `force:",omitempty"`
	Cpu              int    `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Dml              int    `force:",omitempty"`
	DmlRows          int    `force:",omitempty"`
	Email            int    `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	LimitContext     string `force:",omitempty"`
	LimitExceptions  string `force:",omitempty"`
	MobilePush       int    `force:",omitempty"`
	QueryRows        int    `force:",omitempty"`
	Soql             int    `force:",omitempty"`
	Sosl             int    `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *ApexTestResultLimits) ApiName() string {
	return "ApexTestResultLimits"
}

func (t *ApexTestResultLimits) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ApexTestResultLimits #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApexTestResultId: %v\n", t.ApexTestResultId))
	builder.WriteString(fmt.Sprintf("\tAsyncCalls: %v\n", t.AsyncCalls))
	builder.WriteString(fmt.Sprintf("\tCallouts: %v\n", t.Callouts))
	builder.WriteString(fmt.Sprintf("\tCpu: %v\n", t.Cpu))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDml: %v\n", t.Dml))
	builder.WriteString(fmt.Sprintf("\tDmlRows: %v\n", t.DmlRows))
	builder.WriteString(fmt.Sprintf("\tEmail: %v\n", t.Email))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLimitContext: %v\n", t.LimitContext))
	builder.WriteString(fmt.Sprintf("\tLimitExceptions: %v\n", t.LimitExceptions))
	builder.WriteString(fmt.Sprintf("\tMobilePush: %v\n", t.MobilePush))
	builder.WriteString(fmt.Sprintf("\tQueryRows: %v\n", t.QueryRows))
	builder.WriteString(fmt.Sprintf("\tSoql: %v\n", t.Soql))
	builder.WriteString(fmt.Sprintf("\tSosl: %v\n", t.Sosl))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ApexTestResultLimitsQueryResponse struct {
	BaseQuery
	Records []ApexTestResultLimits `json:"Records" force:"records"`
}
