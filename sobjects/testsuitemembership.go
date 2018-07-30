// This file was generated for SObject TestSuiteMembership, API Version v43.0 at 2018-07-30 03:47:25.317928053 -0400 EDT m=+11.660972043

package sobjects

import (
	"fmt"
	"strings"
)

type TestSuiteMembership struct {
	BaseSObject
	ApexClassId      string `force:",omitempty"`
	ApexTestSuiteId  string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *TestSuiteMembership) ApiName() string {
	return "TestSuiteMembership"
}

func (t *TestSuiteMembership) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("TestSuiteMembership #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApexClassId: %v\n", t.ApexClassId))
	builder.WriteString(fmt.Sprintf("\tApexTestSuiteId: %v\n", t.ApexTestSuiteId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type TestSuiteMembershipQueryResponse struct {
	BaseQuery
	Records []TestSuiteMembership `json:"Records" force:"records"`
}
