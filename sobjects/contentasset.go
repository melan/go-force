// This file was generated for SObject ContentAsset, API Version v43.0 at 2018-07-30 03:47:45.265800687 -0400 EDT m=+31.609593200

package sobjects

import (
	"fmt"
	"strings"
)

type ContentAsset struct {
	BaseSObject
	ContentDocumentId        string `force:",omitempty"`
	CreatedById              string `force:",omitempty"`
	CreatedDate              string `force:",omitempty"`
	DeveloperName            string `force:",omitempty"`
	Id                       string `force:",omitempty"`
	IsDeleted                bool   `force:",omitempty"`
	IsVisibleByExternalUsers bool   `force:",omitempty"`
	Language                 string `force:",omitempty"`
	LastModifiedById         string `force:",omitempty"`
	LastModifiedDate         string `force:",omitempty"`
	MasterLabel              string `force:",omitempty"`
	NamespacePrefix          string `force:",omitempty"`
	SystemModstamp           string `force:",omitempty"`
}

func (t *ContentAsset) ApiName() string {
	return "ContentAsset"
}

func (t *ContentAsset) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentAsset #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentDocumentId: %v\n", t.ContentDocumentId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsVisibleByExternalUsers: %v\n", t.IsVisibleByExternalUsers))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ContentAssetQueryResponse struct {
	BaseQuery
	Records []ContentAsset `json:"Records" force:"records"`
}
