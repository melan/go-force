// This file was generated for SObject LightningUsageByAppTypeMetrics, API Version v43.0 at 2018-07-30 03:48:12.368679917 -0400 EDT m=+58.713489438

package sobjects

import (
	"fmt"
	"strings"
)

type LightningUsageByAppTypeMetrics struct {
	BaseSObject
	AppExperience  string `force:",omitempty"`
	Id             string `force:",omitempty"`
	MetricsDate    string `force:",omitempty"`
	SystemModstamp string `force:",omitempty"`
	UserId         string `force:",omitempty"`
}

func (t *LightningUsageByAppTypeMetrics) ApiName() string {
	return "LightningUsageByAppTypeMetrics"
}

func (t *LightningUsageByAppTypeMetrics) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("LightningUsageByAppTypeMetrics #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAppExperience: %v\n", t.AppExperience))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tMetricsDate: %v\n", t.MetricsDate))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type LightningUsageByAppTypeMetricsQueryResponse struct {
	BaseQuery
	Records []LightningUsageByAppTypeMetrics `json:"Records" force:"records"`
}
