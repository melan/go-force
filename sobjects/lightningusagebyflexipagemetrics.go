// This file was generated for SObject LightningUsageByFlexiPageMetrics, API Version v43.0 at 2018-07-30 03:47:51.993835623 -0400 EDT m=+38.337880599

package sobjects

import (
	"fmt"
	"strings"
)

type LightningUsageByFlexiPageMetrics struct {
	BaseSObject
	FlexiPageNameOrId string `force:",omitempty"`
	FlexiPageType     string `force:",omitempty"`
	Id                string `force:",omitempty"`
	MetricsDate       string `force:",omitempty"`
	SystemModstamp    string `force:",omitempty"`
	TotalCount        int    `force:",omitempty"`
}

func (t *LightningUsageByFlexiPageMetrics) ApiName() string {
	return "LightningUsageByFlexiPageMetrics"
}

func (t *LightningUsageByFlexiPageMetrics) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("LightningUsageByFlexiPageMetrics #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tFlexiPageNameOrId: %v\n", t.FlexiPageNameOrId))
	builder.WriteString(fmt.Sprintf("\tFlexiPageType: %v\n", t.FlexiPageType))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tMetricsDate: %v\n", t.MetricsDate))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTotalCount: %v\n", t.TotalCount))

	return builder.String()
}

type LightningUsageByFlexiPageMetricsQueryResponse struct {
	BaseQuery
	Records []LightningUsageByFlexiPageMetrics `json:"Records" force:"records"`
}
