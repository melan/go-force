// This file was generated for SObject ContentUserSubscription, API Version v43.0 at 2018-07-30 03:48:05.652546309 -0400 EDT m=+51.997103814

package sobjects

import (
	"fmt"
	"strings"
)

type ContentUserSubscription struct {
	BaseSObject
	Id                 string `force:",omitempty"`
	SubscribedToUserId string `force:",omitempty"`
	SubscriberUserId   string `force:",omitempty"`
}

func (t *ContentUserSubscription) ApiName() string {
	return "ContentUserSubscription"
}

func (t *ContentUserSubscription) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentUserSubscription #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tSubscribedToUserId: %v\n", t.SubscribedToUserId))
	builder.WriteString(fmt.Sprintf("\tSubscriberUserId: %v\n", t.SubscriberUserId))

	return builder.String()
}

type ContentUserSubscriptionQueryResponse struct {
	BaseQuery
	Records []ContentUserSubscription `json:"Records" force:"records"`
}
