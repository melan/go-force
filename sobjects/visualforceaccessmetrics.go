// This file was generated for SObject VisualforceAccessMetrics, API Version v43.0 at 2018-07-30 03:47:37.25685993 -0400 EDT m=+23.600351916

package sobjects

import (
	"fmt"
	"strings"
)

type VisualforceAccessMetrics struct {
	BaseSObject
	ApexPageId         string `force:",omitempty"`
	DailyPageViewCount int    `force:",omitempty"`
	Id                 string `force:",omitempty"`
	MetricsDate        string `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
}

func (t *VisualforceAccessMetrics) ApiName() string {
	return "VisualforceAccessMetrics"
}

func (t *VisualforceAccessMetrics) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("VisualforceAccessMetrics #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApexPageId: %v\n", t.ApexPageId))
	builder.WriteString(fmt.Sprintf("\tDailyPageViewCount: %v\n", t.DailyPageViewCount))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tMetricsDate: %v\n", t.MetricsDate))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type VisualforceAccessMetricsQueryResponse struct {
	BaseQuery
	Records []VisualforceAccessMetrics `json:"Records" force:"records"`
}
