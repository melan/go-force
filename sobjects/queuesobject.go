// This file was generated for SObject QueueSobject, API Version v43.0 at 2018-07-30 03:47:59.316587951 -0400 EDT m=+45.660907705

package sobjects

import (
	"fmt"
	"strings"
)

type QueueSobject struct {
	BaseSObject
	CreatedById    string `force:",omitempty"`
	Id             string `force:",omitempty"`
	QueueId        string `force:",omitempty"`
	SobjectType    string `force:",omitempty"`
	SystemModstamp string `force:",omitempty"`
}

func (t *QueueSobject) ApiName() string {
	return "QueueSobject"
}

func (t *QueueSobject) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("QueueSobject #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tQueueId: %v\n", t.QueueId))
	builder.WriteString(fmt.Sprintf("\tSobjectType: %v\n", t.SobjectType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type QueueSobjectQueryResponse struct {
	BaseQuery
	Records []QueueSobject `json:"Records" force:"records"`
}
