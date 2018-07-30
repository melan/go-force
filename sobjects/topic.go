// This file was generated for SObject Topic, API Version v43.0 at 2018-07-30 03:47:55.986308731 -0400 EDT m=+42.330503521

package sobjects

import (
	"fmt"
	"strings"
)

type Topic struct {
	BaseSObject
	CreatedById    string `force:",omitempty"`
	CreatedDate    string `force:",omitempty"`
	Description    string `force:",omitempty"`
	Id             string `force:",omitempty"`
	Name           string `force:",omitempty"`
	SystemModstamp string `force:",omitempty"`
	TalkingAbout   int    `force:",omitempty"`
}

func (t *Topic) ApiName() string {
	return "Topic"
}

func (t *Topic) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Topic #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTalkingAbout: %v\n", t.TalkingAbout))

	return builder.String()
}

type TopicQueryResponse struct {
	BaseQuery
	Records []Topic `json:"Records" force:"records"`
}
