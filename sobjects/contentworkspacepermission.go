// This file was generated for SObject ContentWorkspacePermission, API Version v43.0 at 2018-07-30 03:47:45.900435663 -0400 EDT m=+32.244251991

package sobjects

import (
	"fmt"
	"strings"
)

type ContentWorkspacePermission struct {
	BaseSObject
	CreatedById                      string `force:",omitempty"`
	CreatedDate                      string `force:",omitempty"`
	Description                      string `force:",omitempty"`
	Id                               string `force:",omitempty"`
	LastModifiedById                 string `force:",omitempty"`
	LastModifiedDate                 string `force:",omitempty"`
	Name                             string `force:",omitempty"`
	PermissionsAddComment            bool   `force:",omitempty"`
	PermissionsAddContent            bool   `force:",omitempty"`
	PermissionsAddContentOBO         bool   `force:",omitempty"`
	PermissionsArchiveContent        bool   `force:",omitempty"`
	PermissionsChatterSharing        bool   `force:",omitempty"`
	PermissionsDeleteContent         bool   `force:",omitempty"`
	PermissionsDeliverContent        bool   `force:",omitempty"`
	PermissionsFeatureContent        bool   `force:",omitempty"`
	PermissionsManageWorkspace       bool   `force:",omitempty"`
	PermissionsModifyComments        bool   `force:",omitempty"`
	PermissionsOrganizeFileAndFolder bool   `force:",omitempty"`
	PermissionsTagContent            bool   `force:",omitempty"`
	PermissionsViewComments          bool   `force:",omitempty"`
	SystemModstamp                   string `force:",omitempty"`
	Type                             string `force:",omitempty"`
}

func (t *ContentWorkspacePermission) ApiName() string {
	return "ContentWorkspacePermission"
}

func (t *ContentWorkspacePermission) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentWorkspacePermission #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tPermissionsAddComment: %v\n", t.PermissionsAddComment))
	builder.WriteString(fmt.Sprintf("\tPermissionsAddContent: %v\n", t.PermissionsAddContent))
	builder.WriteString(fmt.Sprintf("\tPermissionsAddContentOBO: %v\n", t.PermissionsAddContentOBO))
	builder.WriteString(fmt.Sprintf("\tPermissionsArchiveContent: %v\n", t.PermissionsArchiveContent))
	builder.WriteString(fmt.Sprintf("\tPermissionsChatterSharing: %v\n", t.PermissionsChatterSharing))
	builder.WriteString(fmt.Sprintf("\tPermissionsDeleteContent: %v\n", t.PermissionsDeleteContent))
	builder.WriteString(fmt.Sprintf("\tPermissionsDeliverContent: %v\n", t.PermissionsDeliverContent))
	builder.WriteString(fmt.Sprintf("\tPermissionsFeatureContent: %v\n", t.PermissionsFeatureContent))
	builder.WriteString(fmt.Sprintf("\tPermissionsManageWorkspace: %v\n", t.PermissionsManageWorkspace))
	builder.WriteString(fmt.Sprintf("\tPermissionsModifyComments: %v\n", t.PermissionsModifyComments))
	builder.WriteString(fmt.Sprintf("\tPermissionsOrganizeFileAndFolder: %v\n", t.PermissionsOrganizeFileAndFolder))
	builder.WriteString(fmt.Sprintf("\tPermissionsTagContent: %v\n", t.PermissionsTagContent))
	builder.WriteString(fmt.Sprintf("\tPermissionsViewComments: %v\n", t.PermissionsViewComments))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type ContentWorkspacePermissionQueryResponse struct {
	BaseQuery
	Records []ContentWorkspacePermission `json:"Records" force:"records"`
}
