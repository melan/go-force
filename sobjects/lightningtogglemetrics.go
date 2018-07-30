// This file was generated for SObject LightningToggleMetrics, API Version v43.0 at 2018-07-30 03:47:36.889095059 -0400 EDT m=+23.232573246

package sobjects

import (
	"fmt"
	"strings"
)

type LightningToggleMetrics struct {
	BaseSObject
	Action         string `force:",omitempty"`
	Id             string `force:",omitempty"`
	MetricsDate    string `force:",omitempty"`
	RecordCount    int    `force:",omitempty"`
	SystemModstamp string `force:",omitempty"`
	UserId         string `force:",omitempty"`
}

func (t *LightningToggleMetrics) ApiName() string {
	return "LightningToggleMetrics"
}

func (t *LightningToggleMetrics) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("LightningToggleMetrics #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAction: %v\n", t.Action))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tMetricsDate: %v\n", t.MetricsDate))
	builder.WriteString(fmt.Sprintf("\tRecordCount: %v\n", t.RecordCount))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type LightningToggleMetricsQueryResponse struct {
	BaseQuery
	Records []LightningToggleMetrics `json:"Records" force:"records"`
}
