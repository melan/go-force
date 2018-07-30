// This file was generated for SObject FolderedContentDocument, API Version v43.0 at 2018-07-30 03:48:11.463506983 -0400 EDT m=+57.808282538

package sobjects

import (
	"fmt"
	"strings"
)

type FolderedContentDocument struct {
	BaseSObject
	ContentDocumentId     string `force:",omitempty"`
	ContentSize           int    `force:",omitempty"`
	CreatedById           string `force:",omitempty"`
	CreatedDate           string `force:",omitempty"`
	FileExtension         string `force:",omitempty"`
	FileType              string `force:",omitempty"`
	Id                    string `force:",omitempty"`
	IsDeleted             bool   `force:",omitempty"`
	IsFolder              bool   `force:",omitempty"`
	LastModifiedById      string `force:",omitempty"`
	LastModifiedDate      string `force:",omitempty"`
	ParentContentFolderId string `force:",omitempty"`
	SystemModstamp        string `force:",omitempty"`
	Title                 string `force:",omitempty"`
}

func (t *FolderedContentDocument) ApiName() string {
	return "FolderedContentDocument"
}

func (t *FolderedContentDocument) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FolderedContentDocument #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentDocumentId: %v\n", t.ContentDocumentId))
	builder.WriteString(fmt.Sprintf("\tContentSize: %v\n", t.ContentSize))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFileExtension: %v\n", t.FileExtension))
	builder.WriteString(fmt.Sprintf("\tFileType: %v\n", t.FileType))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsFolder: %v\n", t.IsFolder))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tParentContentFolderId: %v\n", t.ParentContentFolderId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))

	return builder.String()
}

type FolderedContentDocumentQueryResponse struct {
	BaseQuery
	Records []FolderedContentDocument `json:"Records" force:"records"`
}
