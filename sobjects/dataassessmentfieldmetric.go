// This file was generated for SObject DataAssessmentFieldMetric, API Version v43.0 at 2018-07-30 03:47:22.864692461 -0400 EDT m=+9.207644397

package sobjects

import (
	"fmt"
	"strings"
)

type DataAssessmentFieldMetric struct {
	BaseSObject
	CreatedById            string `force:",omitempty"`
	CreatedDate            string `force:",omitempty"`
	DataAssessmentMetricId string `force:",omitempty"`
	FieldName              string `force:",omitempty"`
	Id                     string `force:",omitempty"`
	IsDeleted              bool   `force:",omitempty"`
	LastModifiedById       string `force:",omitempty"`
	LastModifiedDate       string `force:",omitempty"`
	Name                   string `force:",omitempty"`
	NumMatchedBlanks       int    `force:",omitempty"`
	NumMatchedDifferent    int    `force:",omitempty"`
	NumMatchedInSync       int    `force:",omitempty"`
	NumUnmatchedBlanks     int    `force:",omitempty"`
	SystemModstamp         string `force:",omitempty"`
}

func (t *DataAssessmentFieldMetric) ApiName() string {
	return "DataAssessmentFieldMetric"
}

func (t *DataAssessmentFieldMetric) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DataAssessmentFieldMetric #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDataAssessmentMetricId: %v\n", t.DataAssessmentMetricId))
	builder.WriteString(fmt.Sprintf("\tFieldName: %v\n", t.FieldName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNumMatchedBlanks: %v\n", t.NumMatchedBlanks))
	builder.WriteString(fmt.Sprintf("\tNumMatchedDifferent: %v\n", t.NumMatchedDifferent))
	builder.WriteString(fmt.Sprintf("\tNumMatchedInSync: %v\n", t.NumMatchedInSync))
	builder.WriteString(fmt.Sprintf("\tNumUnmatchedBlanks: %v\n", t.NumUnmatchedBlanks))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type DataAssessmentFieldMetricQueryResponse struct {
	BaseQuery
	Records []DataAssessmentFieldMetric `json:"Records" force:"records"`
}
