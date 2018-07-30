// This file was generated for SObject SiteHistory, API Version v43.0 at 2018-07-30 03:47:49.798250672 -0400 EDT m=+36.142213261

package sobjects

import (
	"fmt"
	"strings"
)

type SiteHistory struct {
	BaseSObject
	CreatedById string `force:",omitempty"`
	CreatedDate string `force:",omitempty"`
	Field       string `force:",omitempty"`
	Id          string `force:",omitempty"`
	IsDeleted   bool   `force:",omitempty"`
	NewValue    string `force:",omitempty"`
	OldValue    string `force:",omitempty"`
	SiteId      string `force:",omitempty"`
}

func (t *SiteHistory) ApiName() string {
	return "SiteHistory"
}

func (t *SiteHistory) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("SiteHistory #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))
	builder.WriteString(fmt.Sprintf("\tSiteId: %v\n", t.SiteId))

	return builder.String()
}

type SiteHistoryQueryResponse struct {
	BaseQuery
	Records []SiteHistory `json:"Records" force:"records"`
}
