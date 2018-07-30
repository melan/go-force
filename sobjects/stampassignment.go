// This file was generated for SObject StampAssignment, API Version v43.0 at 2018-07-30 03:48:10.412456583 -0400 EDT m=+56.757192698

package sobjects

import (
	"fmt"
	"strings"
)

type StampAssignment struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	StampId          string `force:",omitempty"`
	SubjectId        string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *StampAssignment) ApiName() string {
	return "StampAssignment"
}

func (t *StampAssignment) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("StampAssignment #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tStampId: %v\n", t.StampId))
	builder.WriteString(fmt.Sprintf("\tSubjectId: %v\n", t.SubjectId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type StampAssignmentQueryResponse struct {
	BaseQuery
	Records []StampAssignment `json:"Records" force:"records"`
}
