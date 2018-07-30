// This file was generated for SObject Announcement, API Version v43.0 at 2018-07-30 03:47:57.821794476 -0400 EDT m=+44.166058140

package sobjects

import (
	"fmt"
	"strings"
)

type Announcement struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	ExpirationDate   string `force:",omitempty"`
	FeedItemId       string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsArchived       bool   `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	ParentId         string `force:",omitempty"`
	SendEmails       bool   `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *Announcement) ApiName() string {
	return "Announcement"
}

func (t *Announcement) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Announcement #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tExpirationDate: %v\n", t.ExpirationDate))
	builder.WriteString(fmt.Sprintf("\tFeedItemId: %v\n", t.FeedItemId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsArchived: %v\n", t.IsArchived))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tSendEmails: %v\n", t.SendEmails))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type AnnouncementQueryResponse struct {
	BaseQuery
	Records []Announcement `json:"Records" force:"records"`
}
