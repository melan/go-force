// This file was generated for SObject ContentNotification, API Version v43.0 at 2018-07-30 03:48:00.510632241 -0400 EDT m=+46.854996801

package sobjects

import (
	"fmt"
	"strings"
)

type ContentNotification struct {
	BaseSObject
	CreatedDate        string `force:",omitempty"`
	EntityIdentifierId string `force:",omitempty"`
	EntityType         string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	Nature             string `force:",omitempty"`
	Subject            string `force:",omitempty"`
	Text               string `force:",omitempty"`
	UsersId            string `force:",omitempty"`
}

func (t *ContentNotification) ApiName() string {
	return "ContentNotification"
}

func (t *ContentNotification) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentNotification #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tEntityIdentifierId: %v\n", t.EntityIdentifierId))
	builder.WriteString(fmt.Sprintf("\tEntityType: %v\n", t.EntityType))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tNature: %v\n", t.Nature))
	builder.WriteString(fmt.Sprintf("\tSubject: %v\n", t.Subject))
	builder.WriteString(fmt.Sprintf("\tText: %v\n", t.Text))
	builder.WriteString(fmt.Sprintf("\tUsersId: %v\n", t.UsersId))

	return builder.String()
}

type ContentNotificationQueryResponse struct {
	BaseQuery
	Records []ContentNotification `json:"Records" force:"records"`
}
