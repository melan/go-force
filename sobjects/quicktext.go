// This file was generated for SObject QuickText, API Version v43.0 at 2018-07-30 03:48:06.219293257 -0400 EDT m=+52.563872028

package sobjects

import (
	"fmt"
	"strings"
)

type QuickText struct {
	BaseSObject
	Category           string `force:",omitempty"`
	Channel            string `force:",omitempty"`
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	LastReferencedDate string `force:",omitempty"`
	LastViewedDate     string `force:",omitempty"`
	Message            string `force:",omitempty"`
	Name               string `force:",omitempty"`
	OwnerId            string `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
}

func (t *QuickText) ApiName() string {
	return "QuickText"
}

func (t *QuickText) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("QuickText #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCategory: %v\n", t.Category))
	builder.WriteString(fmt.Sprintf("\tChannel: %v\n", t.Channel))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tMessage: %v\n", t.Message))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type QuickTextQueryResponse struct {
	BaseQuery
	Records []QuickText `json:"Records" force:"records"`
}
