// This file was generated for SObject CustomObjectUserLicenseMetrics, API Version v43.0 at 2018-07-30 03:47:55.619902082 -0400 EDT m=+41.964083122

package sobjects

import (
	"fmt"
	"strings"
)

type CustomObjectUserLicenseMetrics struct {
	BaseSObject
	CustomObjectId   string `force:",omitempty"`
	CustomObjectName string `force:",omitempty"`
	CustomObjectType string `force:",omitempty"`
	Id               string `force:",omitempty"`
	MetricsDate      string `force:",omitempty"`
	ObjectCount      int    `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	UserLicenseId    string `force:",omitempty"`
}

func (t *CustomObjectUserLicenseMetrics) ApiName() string {
	return "CustomObjectUserLicenseMetrics"
}

func (t *CustomObjectUserLicenseMetrics) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CustomObjectUserLicenseMetrics #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCustomObjectId: %v\n", t.CustomObjectId))
	builder.WriteString(fmt.Sprintf("\tCustomObjectName: %v\n", t.CustomObjectName))
	builder.WriteString(fmt.Sprintf("\tCustomObjectType: %v\n", t.CustomObjectType))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tMetricsDate: %v\n", t.MetricsDate))
	builder.WriteString(fmt.Sprintf("\tObjectCount: %v\n", t.ObjectCount))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserLicenseId: %v\n", t.UserLicenseId))

	return builder.String()
}

type CustomObjectUserLicenseMetricsQueryResponse struct {
	BaseQuery
	Records []CustomObjectUserLicenseMetrics `json:"Records" force:"records"`
}
