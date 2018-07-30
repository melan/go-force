// This file was generated for SObject DataStatistics, API Version v43.0 at 2018-07-30 03:47:35.334347579 -0400 EDT m=+21.677767425

package sobjects

import (
	"fmt"
	"strings"
)

type DataStatistics struct {
	BaseSObject
	ExternalId string `force:",omitempty"`
	Id         string `force:",omitempty"`
	StatType   string `force:",omitempty"`
	StatValue  int    `force:",omitempty"`
	Type       string `force:",omitempty"`
	UserId     string `force:",omitempty"`
}

func (t *DataStatistics) ApiName() string {
	return "DataStatistics"
}

func (t *DataStatistics) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DataStatistics #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tExternalId: %v\n", t.ExternalId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tStatType: %v\n", t.StatType))
	builder.WriteString(fmt.Sprintf("\tStatValue: %v\n", t.StatValue))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type DataStatisticsQueryResponse struct {
	BaseQuery
	Records []DataStatistics `json:"Records" force:"records"`
}
