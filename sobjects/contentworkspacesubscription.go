// This file was generated for SObject ContentWorkspaceSubscription, API Version v43.0 at 2018-07-30 03:47:44.370316549 -0400 EDT m=+30.714075460

package sobjects

import (
	"fmt"
	"strings"
)

type ContentWorkspaceSubscription struct {
	BaseSObject
	ContentWorkspaceId string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	UserId             string `force:",omitempty"`
}

func (t *ContentWorkspaceSubscription) ApiName() string {
	return "ContentWorkspaceSubscription"
}

func (t *ContentWorkspaceSubscription) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentWorkspaceSubscription #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentWorkspaceId: %v\n", t.ContentWorkspaceId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type ContentWorkspaceSubscriptionQueryResponse struct {
	BaseQuery
	Records []ContentWorkspaceSubscription `json:"Records" force:"records"`
}
