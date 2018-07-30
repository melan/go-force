// This file was generated for SObject Group, API Version v43.0 at 2018-07-30 03:47:18.607931507 -0400 EDT m=+4.950723710

package sobjects

import (
	"fmt"
	"strings"
)

type Group struct {
	BaseSObject
	CreatedById            string `force:",omitempty"`
	CreatedDate            string `force:",omitempty"`
	DeveloperName          string `force:",omitempty"`
	DoesIncludeBosses      bool   `force:",omitempty"`
	DoesSendEmailToMembers bool   `force:",omitempty"`
	Email                  string `force:",omitempty"`
	Id                     string `force:",omitempty"`
	LastModifiedById       string `force:",omitempty"`
	LastModifiedDate       string `force:",omitempty"`
	Name                   string `force:",omitempty"`
	OwnerId                string `force:",omitempty"`
	RelatedId              string `force:",omitempty"`
	SystemModstamp         string `force:",omitempty"`
	Type                   string `force:",omitempty"`
}

func (t *Group) ApiName() string {
	return "Group"
}

func (t *Group) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Group #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tDoesIncludeBosses: %v\n", t.DoesIncludeBosses))
	builder.WriteString(fmt.Sprintf("\tDoesSendEmailToMembers: %v\n", t.DoesSendEmailToMembers))
	builder.WriteString(fmt.Sprintf("\tEmail: %v\n", t.Email))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tRelatedId: %v\n", t.RelatedId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type GroupQueryResponse struct {
	BaseQuery
	Records []Group `json:"Records" force:"records"`
}
