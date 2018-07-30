// This file was generated for SObject LightningUsageByPageMetrics, API Version v43.0 at 2018-07-30 03:47:17.340680714 -0400 EDT m=+3.683425364

package sobjects

import (
	"fmt"
	"strings"
)

type LightningUsageByPageMetrics struct {
	BaseSObject
	Id             string `force:",omitempty"`
	MetricsDate    string `force:",omitempty"`
	PageName       string `force:",omitempty"`
	SystemModstamp string `force:",omitempty"`
	TotalCount     int    `force:",omitempty"`
	UserId         string `force:",omitempty"`
}

func (t *LightningUsageByPageMetrics) ApiName() string {
	return "LightningUsageByPageMetrics"
}

func (t *LightningUsageByPageMetrics) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("LightningUsageByPageMetrics #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tMetricsDate: %v\n", t.MetricsDate))
	builder.WriteString(fmt.Sprintf("\tPageName: %v\n", t.PageName))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTotalCount: %v\n", t.TotalCount))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type LightningUsageByPageMetricsQueryResponse struct {
	BaseQuery
	Records []LightningUsageByPageMetrics `json:"Records" force:"records"`
}
