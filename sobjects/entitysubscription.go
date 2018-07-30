// This file was generated for SObject EntitySubscription, API Version v43.0 at 2018-07-30 03:47:49.09697427 -0400 EDT m=+35.440910545

package sobjects

import (
	"fmt"
	"strings"
)

type EntitySubscription struct {
	BaseSObject
	CreatedById  string `force:",omitempty"`
	CreatedDate  string `force:",omitempty"`
	Id           string `force:",omitempty"`
	IsDeleted    bool   `force:",omitempty"`
	ParentId     string `force:",omitempty"`
	SubscriberId string `force:",omitempty"`
}

func (t *EntitySubscription) ApiName() string {
	return "EntitySubscription"
}

func (t *EntitySubscription) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("EntitySubscription #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tSubscriberId: %v\n", t.SubscriberId))

	return builder.String()
}

type EntitySubscriptionQueryResponse struct {
	BaseQuery
	Records []EntitySubscription `json:"Records" force:"records"`
}
