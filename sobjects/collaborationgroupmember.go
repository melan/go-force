// This file was generated for SObject CollaborationGroupMember, API Version v43.0 at 2018-07-30 03:47:52.935115653 -0400 EDT m=+39.279195950

package sobjects

import (
	"fmt"
	"strings"
)

type CollaborationGroupMember struct {
	BaseSObject
	CollaborationGroupId  string `force:",omitempty"`
	CollaborationRole     string `force:",omitempty"`
	CreatedById           string `force:",omitempty"`
	CreatedDate           string `force:",omitempty"`
	Id                    string `force:",omitempty"`
	LastFeedAccessDate    string `force:",omitempty"`
	LastModifiedById      string `force:",omitempty"`
	LastModifiedDate      string `force:",omitempty"`
	MemberId              string `force:",omitempty"`
	NotificationFrequency string `force:",omitempty"`
	SystemModstamp        string `force:",omitempty"`
}

func (t *CollaborationGroupMember) ApiName() string {
	return "CollaborationGroupMember"
}

func (t *CollaborationGroupMember) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CollaborationGroupMember #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCollaborationGroupId: %v\n", t.CollaborationGroupId))
	builder.WriteString(fmt.Sprintf("\tCollaborationRole: %v\n", t.CollaborationRole))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastFeedAccessDate: %v\n", t.LastFeedAccessDate))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMemberId: %v\n", t.MemberId))
	builder.WriteString(fmt.Sprintf("\tNotificationFrequency: %v\n", t.NotificationFrequency))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CollaborationGroupMemberQueryResponse struct {
	BaseQuery
	Records []CollaborationGroupMember `json:"Records" force:"records"`
}
