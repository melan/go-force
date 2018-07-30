// This file was generated for SObject EventBusSubscriber, API Version v43.0 at 2018-07-30 03:47:55.497737848 -0400 EDT m=+41.841914304

package sobjects

import (
	"fmt"
	"strings"
)

type EventBusSubscriber struct {
	BaseSObject
	ExternalId string `force:",omitempty"`
	Id         string `force:",omitempty"`
	LastError  string `force:",omitempty"`
	Name       string `force:",omitempty"`
	Position   int    `force:",omitempty"`
	Retries    int    `force:",omitempty"`
	Status     string `force:",omitempty"`
	Tip        int    `force:",omitempty"`
	Topic      string `force:",omitempty"`
	Type       string `force:",omitempty"`
}

func (t *EventBusSubscriber) ApiName() string {
	return "EventBusSubscriber"
}

func (t *EventBusSubscriber) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("EventBusSubscriber #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tExternalId: %v\n", t.ExternalId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastError: %v\n", t.LastError))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tPosition: %v\n", t.Position))
	builder.WriteString(fmt.Sprintf("\tRetries: %v\n", t.Retries))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tTip: %v\n", t.Tip))
	builder.WriteString(fmt.Sprintf("\tTopic: %v\n", t.Topic))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type EventBusSubscriberQueryResponse struct {
	BaseQuery
	Records []EventBusSubscriber `json:"Records" force:"records"`
}
