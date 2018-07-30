// This file was generated for SObject ContentWorkspace, API Version v43.0 at 2018-07-30 03:47:20.342433208 -0400 EDT m=+6.685290497

package sobjects

import (
	"fmt"
	"strings"
)

type ContentWorkspace struct {
	BaseSObject
	CreatedById                  string `force:",omitempty"`
	CreatedDate                  string `force:",omitempty"`
	DefaultRecordTypeId          string `force:",omitempty"`
	Description                  string `force:",omitempty"`
	DeveloperName                string `force:",omitempty"`
	Id                           string `force:",omitempty"`
	IsRestrictContentTypes       bool   `force:",omitempty"`
	IsRestrictLinkedContentTypes bool   `force:",omitempty"`
	LastModifiedById             string `force:",omitempty"`
	LastModifiedDate             string `force:",omitempty"`
	LastWorkspaceActivityDate    string `force:",omitempty"`
	Name                         string `force:",omitempty"`
	NamespacePrefix              string `force:",omitempty"`
	RootContentFolderId          string `force:",omitempty"`
	ShouldAddCreatorMembership   bool   `force:",omitempty"`
	SystemModstamp               string `force:",omitempty"`
	TagModel                     string `force:",omitempty"`
	WorkspaceImageId             string `force:",omitempty"`
	WorkspaceType                string `force:",omitempty"`
}

func (t *ContentWorkspace) ApiName() string {
	return "ContentWorkspace"
}

func (t *ContentWorkspace) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentWorkspace #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDefaultRecordTypeId: %v\n", t.DefaultRecordTypeId))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsRestrictContentTypes: %v\n", t.IsRestrictContentTypes))
	builder.WriteString(fmt.Sprintf("\tIsRestrictLinkedContentTypes: %v\n", t.IsRestrictLinkedContentTypes))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastWorkspaceActivityDate: %v\n", t.LastWorkspaceActivityDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tRootContentFolderId: %v\n", t.RootContentFolderId))
	builder.WriteString(fmt.Sprintf("\tShouldAddCreatorMembership: %v\n", t.ShouldAddCreatorMembership))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTagModel: %v\n", t.TagModel))
	builder.WriteString(fmt.Sprintf("\tWorkspaceImageId: %v\n", t.WorkspaceImageId))
	builder.WriteString(fmt.Sprintf("\tWorkspaceType: %v\n", t.WorkspaceType))

	return builder.String()
}

type ContentWorkspaceQueryResponse struct {
	BaseQuery
	Records []ContentWorkspace `json:"Records" force:"records"`
}
