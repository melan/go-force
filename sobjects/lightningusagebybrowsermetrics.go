// This file was generated for SObject LightningUsageByBrowserMetrics, API Version v43.0 at 2018-07-30 03:47:25.432567368 -0400 EDT m=+11.775615660

package sobjects

import (
	"fmt"
	"strings"
)

type LightningUsageByBrowserMetrics struct {
	BaseSObject
	Browser        string `force:",omitempty"`
	Id             string `force:",omitempty"`
	MetricsDate    string `force:",omitempty"`
	PageName       string `force:",omitempty"`
	SystemModstamp string `force:",omitempty"`
	TotalCount     int    `force:",omitempty"`
}

func (t *LightningUsageByBrowserMetrics) ApiName() string {
	return "LightningUsageByBrowserMetrics"
}

func (t *LightningUsageByBrowserMetrics) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("LightningUsageByBrowserMetrics #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tBrowser: %v\n", t.Browser))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tMetricsDate: %v\n", t.MetricsDate))
	builder.WriteString(fmt.Sprintf("\tPageName: %v\n", t.PageName))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTotalCount: %v\n", t.TotalCount))

	return builder.String()
}

type LightningUsageByBrowserMetricsQueryResponse struct {
	BaseQuery
	Records []LightningUsageByBrowserMetrics `json:"Records" force:"records"`
}
