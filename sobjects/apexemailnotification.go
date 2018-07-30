// This file was generated for SObject ApexEmailNotification, API Version v43.0 at 2018-07-30 03:47:37.514094028 -0400 EDT m=+23.857595667

package sobjects

import (
	"fmt"
	"strings"
)

type ApexEmailNotification struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Email            string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	UserId           string `force:",omitempty"`
}

func (t *ApexEmailNotification) ApiName() string {
	return "ApexEmailNotification"
}

func (t *ApexEmailNotification) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ApexEmailNotification #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tEmail: %v\n", t.Email))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type ApexEmailNotificationQueryResponse struct {
	BaseQuery
	Records []ApexEmailNotification `json:"Records" force:"records"`
}
