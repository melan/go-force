// This file was generated for SObject ApexTestSuite, API Version v43.0 at 2018-07-30 03:47:19.814166884 -0400 EDT m=+6.157004350

package sobjects

import (
	"fmt"
	"strings"
)

type ApexTestSuite struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	TestSuiteName    string `force:",omitempty"`
}

func (t *ApexTestSuite) ApiName() string {
	return "ApexTestSuite"
}

func (t *ApexTestSuite) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ApexTestSuite #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTestSuiteName: %v\n", t.TestSuiteName))

	return builder.String()
}

type ApexTestSuiteQueryResponse struct {
	BaseQuery
	Records []ApexTestSuite `json:"Records" force:"records"`
}
