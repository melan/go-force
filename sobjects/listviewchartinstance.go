// This file was generated for SObject ListViewChartInstance, API Version v43.0 at 2018-07-30 03:47:55.087705267 -0400 EDT m=+41.431866338

package sobjects

import (
	"fmt"
	"strings"
)

type ListViewChartInstance struct {
	BaseSObject
	AggregateField              string `force:",omitempty"`
	AggregateType               string `force:",omitempty"`
	ChartType                   string `force:",omitempty"`
	DataQuery                   string `force:",omitempty"`
	DataQueryWithoutUserFilters string `force:",omitempty"`
	DeveloperName               string `force:",omitempty"`
	ExternalId                  string `force:",omitempty"`
	GroupingField               string `force:",omitempty"`
	Id                          string `force:",omitempty"`
	IsDeletable                 bool   `force:",omitempty"`
	IsEditable                  bool   `force:",omitempty"`
	IsLastViewed                bool   `force:",omitempty"`
	Label                       string `force:",omitempty"`
	ListViewChartId             string `force:",omitempty"`
	ListViewContextId           string `force:",omitempty"`
	SourceEntity                string `force:",omitempty"`
}

func (t *ListViewChartInstance) ApiName() string {
	return "ListViewChartInstance"
}

func (t *ListViewChartInstance) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ListViewChartInstance #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAggregateField: %v\n", t.AggregateField))
	builder.WriteString(fmt.Sprintf("\tAggregateType: %v\n", t.AggregateType))
	builder.WriteString(fmt.Sprintf("\tChartType: %v\n", t.ChartType))
	builder.WriteString(fmt.Sprintf("\tDataQuery: %v\n", t.DataQuery))
	builder.WriteString(fmt.Sprintf("\tDataQueryWithoutUserFilters: %v\n", t.DataQueryWithoutUserFilters))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tExternalId: %v\n", t.ExternalId))
	builder.WriteString(fmt.Sprintf("\tGroupingField: %v\n", t.GroupingField))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeletable: %v\n", t.IsDeletable))
	builder.WriteString(fmt.Sprintf("\tIsEditable: %v\n", t.IsEditable))
	builder.WriteString(fmt.Sprintf("\tIsLastViewed: %v\n", t.IsLastViewed))
	builder.WriteString(fmt.Sprintf("\tLabel: %v\n", t.Label))
	builder.WriteString(fmt.Sprintf("\tListViewChartId: %v\n", t.ListViewChartId))
	builder.WriteString(fmt.Sprintf("\tListViewContextId: %v\n", t.ListViewContextId))
	builder.WriteString(fmt.Sprintf("\tSourceEntity: %v\n", t.SourceEntity))

	return builder.String()
}

type ListViewChartInstanceQueryResponse struct {
	BaseQuery
	Records []ListViewChartInstance `json:"Records" force:"records"`
}
