// This file was generated for SObject TopicAssignment, API Version v43.0 at 2018-07-30 03:47:28.532269043 -0400 EDT m=+14.875433648

package sobjects

import (
	"fmt"
	"strings"
)

type TopicAssignment struct {
	BaseSObject
	CreatedById     string `force:",omitempty"`
	CreatedDate     string `force:",omitempty"`
	EntityId        string `force:",omitempty"`
	EntityKeyPrefix string `force:",omitempty"`
	EntityType      string `force:",omitempty"`
	Id              string `force:",omitempty"`
	IsDeleted       bool   `force:",omitempty"`
	SystemModstamp  string `force:",omitempty"`
	TopicId         string `force:",omitempty"`
}

func (t *TopicAssignment) ApiName() string {
	return "TopicAssignment"
}

func (t *TopicAssignment) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("TopicAssignment #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tEntityId: %v\n", t.EntityId))
	builder.WriteString(fmt.Sprintf("\tEntityKeyPrefix: %v\n", t.EntityKeyPrefix))
	builder.WriteString(fmt.Sprintf("\tEntityType: %v\n", t.EntityType))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTopicId: %v\n", t.TopicId))

	return builder.String()
}

type TopicAssignmentQueryResponse struct {
	BaseQuery
	Records []TopicAssignment `json:"Records" force:"records"`
}
