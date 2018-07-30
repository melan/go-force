// This file was generated for SObject FlexQueueItem, API Version v43.0 at 2018-07-30 03:48:03.014735922 -0400 EDT m=+49.359194446

package sobjects

import (
	"fmt"
	"strings"
)

type FlexQueueItem struct {
	BaseSObject
	AsyncApexJobId  string `force:",omitempty"`
	FlexQueueItemId string `force:",omitempty"`
	Id              string `force:",omitempty"`
	JobPosition     int    `force:",omitempty"`
	JobType         string `force:",omitempty"`
}

func (t *FlexQueueItem) ApiName() string {
	return "FlexQueueItem"
}

func (t *FlexQueueItem) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FlexQueueItem #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAsyncApexJobId: %v\n", t.AsyncApexJobId))
	builder.WriteString(fmt.Sprintf("\tFlexQueueItemId: %v\n", t.FlexQueueItemId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tJobPosition: %v\n", t.JobPosition))
	builder.WriteString(fmt.Sprintf("\tJobType: %v\n", t.JobType))

	return builder.String()
}

type FlexQueueItemQueryResponse struct {
	BaseQuery
	Records []FlexQueueItem `json:"Records" force:"records"`
}
