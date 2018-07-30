// This file was generated for SObject DataAssessmentMetric, API Version v43.0 at 2018-07-30 03:47:24.804444664 -0400 EDT m=+11.147469387

package sobjects

import (
	"fmt"
	"strings"
)

type DataAssessmentMetric struct {
	BaseSObject
	CreatedById         string `force:",omitempty"`
	CreatedDate         string `force:",omitempty"`
	Id                  string `force:",omitempty"`
	IsDeleted           bool   `force:",omitempty"`
	LastModifiedById    string `force:",omitempty"`
	LastModifiedDate    string `force:",omitempty"`
	Name                string `force:",omitempty"`
	NumDuplicates       int    `force:",omitempty"`
	NumMatched          int    `force:",omitempty"`
	NumMatchedDifferent int    `force:",omitempty"`
	NumProcessed        int    `force:",omitempty"`
	NumTotal            int    `force:",omitempty"`
	NumUnmatched        int    `force:",omitempty"`
	SystemModstamp      string `force:",omitempty"`
}

func (t *DataAssessmentMetric) ApiName() string {
	return "DataAssessmentMetric"
}

func (t *DataAssessmentMetric) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DataAssessmentMetric #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNumDuplicates: %v\n", t.NumDuplicates))
	builder.WriteString(fmt.Sprintf("\tNumMatched: %v\n", t.NumMatched))
	builder.WriteString(fmt.Sprintf("\tNumMatchedDifferent: %v\n", t.NumMatchedDifferent))
	builder.WriteString(fmt.Sprintf("\tNumProcessed: %v\n", t.NumProcessed))
	builder.WriteString(fmt.Sprintf("\tNumTotal: %v\n", t.NumTotal))
	builder.WriteString(fmt.Sprintf("\tNumUnmatched: %v\n", t.NumUnmatched))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type DataAssessmentMetricQueryResponse struct {
	BaseQuery
	Records []DataAssessmentMetric `json:"Records" force:"records"`
}
