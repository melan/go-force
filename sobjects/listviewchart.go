// This file was generated for SObject ListViewChart, API Version v43.0 at 2018-07-30 03:48:09.776494328 -0400 EDT m=+56.121206580

package sobjects

import (
	"fmt"
	"strings"
)

type ListViewChart struct {
	BaseSObject
	AggregateField   string `force:",omitempty"`
	AggregateType    string `force:",omitempty"`
	ChartType        string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	DeveloperName    string `force:",omitempty"`
	GroupingField    string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	Language         string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	MasterLabel      string `force:",omitempty"`
	OwnerId          string `force:",omitempty"`
	SobjectType      string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *ListViewChart) ApiName() string {
	return "ListViewChart"
}

func (t *ListViewChart) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ListViewChart #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAggregateField: %v\n", t.AggregateField))
	builder.WriteString(fmt.Sprintf("\tAggregateType: %v\n", t.AggregateType))
	builder.WriteString(fmt.Sprintf("\tChartType: %v\n", t.ChartType))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tGroupingField: %v\n", t.GroupingField))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tSobjectType: %v\n", t.SobjectType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ListViewChartQueryResponse struct {
	BaseQuery
	Records []ListViewChart `json:"Records" force:"records"`
}
