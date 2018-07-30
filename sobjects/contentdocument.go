// This file was generated for SObject ContentDocument, API Version v43.0 at 2018-07-30 03:47:59.979096702 -0400 EDT m=+46.323441317

package sobjects

import (
	"fmt"
	"strings"
)

type ContentDocument struct {
	BaseSObject
	ArchivedById             string `force:",omitempty"`
	ArchivedDate             string `force:",omitempty"`
	ContentAssetId           string `force:",omitempty"`
	ContentModifiedDate      string `force:",omitempty"`
	ContentSize              int    `force:",omitempty"`
	CreatedById              string `force:",omitempty"`
	CreatedDate              string `force:",omitempty"`
	Description              string `force:",omitempty"`
	FileExtension            string `force:",omitempty"`
	FileType                 string `force:",omitempty"`
	Id                       string `force:",omitempty"`
	IsArchived               bool   `force:",omitempty"`
	IsDeleted                bool   `force:",omitempty"`
	LastModifiedById         string `force:",omitempty"`
	LastModifiedDate         string `force:",omitempty"`
	LastReferencedDate       string `force:",omitempty"`
	LastViewedDate           string `force:",omitempty"`
	LatestPublishedVersionId string `force:",omitempty"`
	OwnerId                  string `force:",omitempty"`
	ParentId                 string `force:",omitempty"`
	PublishStatus            string `force:",omitempty"`
	SharingOption            string `force:",omitempty"`
	SharingPrivacy           string `force:",omitempty"`
	SystemModstamp           string `force:",omitempty"`
	Title                    string `force:",omitempty"`
}

func (t *ContentDocument) ApiName() string {
	return "ContentDocument"
}

func (t *ContentDocument) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentDocument #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tArchivedById: %v\n", t.ArchivedById))
	builder.WriteString(fmt.Sprintf("\tArchivedDate: %v\n", t.ArchivedDate))
	builder.WriteString(fmt.Sprintf("\tContentAssetId: %v\n", t.ContentAssetId))
	builder.WriteString(fmt.Sprintf("\tContentModifiedDate: %v\n", t.ContentModifiedDate))
	builder.WriteString(fmt.Sprintf("\tContentSize: %v\n", t.ContentSize))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tFileExtension: %v\n", t.FileExtension))
	builder.WriteString(fmt.Sprintf("\tFileType: %v\n", t.FileType))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsArchived: %v\n", t.IsArchived))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tLatestPublishedVersionId: %v\n", t.LatestPublishedVersionId))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tPublishStatus: %v\n", t.PublishStatus))
	builder.WriteString(fmt.Sprintf("\tSharingOption: %v\n", t.SharingOption))
	builder.WriteString(fmt.Sprintf("\tSharingPrivacy: %v\n", t.SharingPrivacy))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))

	return builder.String()
}

type ContentDocumentQueryResponse struct {
	BaseQuery
	Records []ContentDocument `json:"Records" force:"records"`
}
