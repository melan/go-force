// This file was generated for SObject KnowledgeableUser, API Version v43.0 at 2018-07-30 03:47:26.0731645 -0400 EDT m=+12.416236830

package sobjects

import (
	"fmt"
	"strings"
)

type KnowledgeableUser struct {
	BaseSObject
	Id             string `force:",omitempty"`
	RawRank        int    `force:",omitempty"`
	SystemModstamp string `force:",omitempty"`
	TopicId        string `force:",omitempty"`
	UserId         string `force:",omitempty"`
}

func (t *KnowledgeableUser) ApiName() string {
	return "KnowledgeableUser"
}

func (t *KnowledgeableUser) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("KnowledgeableUser #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tRawRank: %v\n", t.RawRank))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTopicId: %v\n", t.TopicId))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type KnowledgeableUserQueryResponse struct {
	BaseQuery
	Records []KnowledgeableUser `json:"Records" force:"records"`
}
