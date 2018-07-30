// This file was generated for SObject RecentlyViewed, API Version v43.0 at 2018-07-30 03:47:17.919045981 -0400 EDT m=+4.261812334

package sobjects

import (
	"fmt"
	"strings"
)

type RecentlyViewed struct {
	BaseSObject
	Alias              string `force:",omitempty"`
	Email              string `force:",omitempty"`
	FirstName          string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsActive           bool   `force:",omitempty"`
	Language           string `force:",omitempty"`
	LastName           string `force:",omitempty"`
	LastReferencedDate string `force:",omitempty"`
	LastViewedDate     string `force:",omitempty"`
	Name               string `force:",omitempty"`
	NameOrAlias        string `force:",omitempty"`
	Phone              string `force:",omitempty"`
	ProfileId          string `force:",omitempty"`
	RecordTypeId       string `force:",omitempty"`
	Title              string `force:",omitempty"`
	Type               string `force:",omitempty"`
	UserRoleId         string `force:",omitempty"`
}

func (t *RecentlyViewed) ApiName() string {
	return "RecentlyViewed"
}

func (t *RecentlyViewed) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("RecentlyViewed #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAlias: %v\n", t.Alias))
	builder.WriteString(fmt.Sprintf("\tEmail: %v\n", t.Email))
	builder.WriteString(fmt.Sprintf("\tFirstName: %v\n", t.FirstName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastName: %v\n", t.LastName))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNameOrAlias: %v\n", t.NameOrAlias))
	builder.WriteString(fmt.Sprintf("\tPhone: %v\n", t.Phone))
	builder.WriteString(fmt.Sprintf("\tProfileId: %v\n", t.ProfileId))
	builder.WriteString(fmt.Sprintf("\tRecordTypeId: %v\n", t.RecordTypeId))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))
	builder.WriteString(fmt.Sprintf("\tUserRoleId: %v\n", t.UserRoleId))

	return builder.String()
}

type RecentlyViewedQueryResponse struct {
	BaseQuery
	Records []RecentlyViewed `json:"Records" force:"records"`
}
