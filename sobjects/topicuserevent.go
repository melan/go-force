// This file was generated for SObject TopicUserEvent, API Version v43.0 at 2018-07-30 03:47:51.127432943 -0400 EDT m=+37.471445409

package sobjects

import (
	"fmt"
	"strings"
)

type TopicUserEvent struct {
	BaseSObject
	ActionEnum  string `force:",omitempty"`
	CreatedDate string `force:",omitempty"`
	Id          string `force:",omitempty"`
	TopicId     string `force:",omitempty"`
	UserId      string `force:",omitempty"`
}

func (t *TopicUserEvent) ApiName() string {
	return "TopicUserEvent"
}

func (t *TopicUserEvent) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("TopicUserEvent #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tActionEnum: %v\n", t.ActionEnum))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tTopicId: %v\n", t.TopicId))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type TopicUserEventQueryResponse struct {
	BaseQuery
	Records []TopicUserEvent `json:"Records" force:"records"`
}
