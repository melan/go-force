// This file was generated for SObject ContentTagSubscription, API Version v43.0 at 2018-07-30 03:47:24.185545877 -0400 EDT m=+10.528547376

package sobjects

import (
	"fmt"
	"strings"
)

type ContentTagSubscription struct {
	BaseSObject
	Id     string `force:",omitempty"`
	UserId string `force:",omitempty"`
}

func (t *ContentTagSubscription) ApiName() string {
	return "ContentTagSubscription"
}

func (t *ContentTagSubscription) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentTagSubscription #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type ContentTagSubscriptionQueryResponse struct {
	BaseQuery
	Records []ContentTagSubscription `json:"Records" force:"records"`
}
