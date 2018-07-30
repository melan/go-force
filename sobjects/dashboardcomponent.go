// This file was generated for SObject DashboardComponent, API Version v43.0 at 2018-07-30 03:47:21.771694116 -0400 EDT m=+8.114605037

package sobjects

import (
	"fmt"
	"strings"
)

type DashboardComponent struct {
	BaseSObject
	CustomReportId string `force:",omitempty"`
	DashboardId    string `force:",omitempty"`
	Id             string `force:",omitempty"`
	Name           string `force:",omitempty"`
}

func (t *DashboardComponent) ApiName() string {
	return "DashboardComponent"
}

func (t *DashboardComponent) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DashboardComponent #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCustomReportId: %v\n", t.CustomReportId))
	builder.WriteString(fmt.Sprintf("\tDashboardId: %v\n", t.DashboardId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))

	return builder.String()
}

type DashboardComponentQueryResponse struct {
	BaseQuery
	Records []DashboardComponent `json:"Records" force:"records"`
}
