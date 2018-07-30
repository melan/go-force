// This file was generated for SObject ListEmail, API Version v43.0 at 2018-07-30 03:47:55.229522601 -0400 EDT m=+41.573688993

package sobjects

import (
	"fmt"
	"strings"
)

type ListEmail struct {
	BaseSObject
	CampaignId         string `force:",omitempty"`
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	FromAddress        string `force:",omitempty"`
	FromName           string `force:",omitempty"`
	HasAttachment      bool   `force:",omitempty"`
	HtmlBody           string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	LastReferencedDate string `force:",omitempty"`
	LastViewedDate     string `force:",omitempty"`
	Name               string `force:",omitempty"`
	OwnerId            string `force:",omitempty"`
	ScheduledDate      string `force:",omitempty"`
	Status             string `force:",omitempty"`
	Subject            string `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
	TextBody           string `force:",omitempty"`
	TotalSent          int    `force:",omitempty"`
}

func (t *ListEmail) ApiName() string {
	return "ListEmail"
}

func (t *ListEmail) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ListEmail #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCampaignId: %v\n", t.CampaignId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFromAddress: %v\n", t.FromAddress))
	builder.WriteString(fmt.Sprintf("\tFromName: %v\n", t.FromName))
	builder.WriteString(fmt.Sprintf("\tHasAttachment: %v\n", t.HasAttachment))
	builder.WriteString(fmt.Sprintf("\tHtmlBody: %v\n", t.HtmlBody))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tScheduledDate: %v\n", t.ScheduledDate))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSubject: %v\n", t.Subject))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTextBody: %v\n", t.TextBody))
	builder.WriteString(fmt.Sprintf("\tTotalSent: %v\n", t.TotalSent))

	return builder.String()
}

type ListEmailQueryResponse struct {
	BaseQuery
	Records []ListEmail `json:"Records" force:"records"`
}
