// This file was generated for SObject DataAssessmentValueMetric, API Version v43.0 at 2018-07-30 03:47:28.917122619 -0400 EDT m=+15.260301666

package sobjects

import (
	"fmt"
	"strings"
)

type DataAssessmentValueMetric struct {
	BaseSObject
	CreatedById                 string `force:",omitempty"`
	CreatedDate                 string `force:",omitempty"`
	DataAssessmentFieldMetricId string `force:",omitempty"`
	FieldValue                  string `force:",omitempty"`
	Id                          string `force:",omitempty"`
	IsDeleted                   bool   `force:",omitempty"`
	LastModifiedById            string `force:",omitempty"`
	LastModifiedDate            string `force:",omitempty"`
	Name                        string `force:",omitempty"`
	SystemModstamp              string `force:",omitempty"`
	ValueCount                  int    `force:",omitempty"`
}

func (t *DataAssessmentValueMetric) ApiName() string {
	return "DataAssessmentValueMetric"
}

func (t *DataAssessmentValueMetric) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DataAssessmentValueMetric #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDataAssessmentFieldMetricId: %v\n", t.DataAssessmentFieldMetricId))
	builder.WriteString(fmt.Sprintf("\tFieldValue: %v\n", t.FieldValue))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tValueCount: %v\n", t.ValueCount))

	return builder.String()
}

type DataAssessmentValueMetricQueryResponse struct {
	BaseQuery
	Records []DataAssessmentValueMetric `json:"Records" force:"records"`
}
